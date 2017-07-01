package bothandlers

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/mgerb/go-discord-bot/server/config"
	"layeh.com/gopus"
)

const (
	channels  int = 2                   // 1 for mono, 2 for stereo
	frameRate int = 48000               // audio sampling rate
	frameSize int = 960                 // uint16 size of each audio frame
	maxBytes  int = (frameSize * 2) * 2 // max size of opus data
)

// store our connection objects in a map tied to a guild id
var activeConnections = make(map[string]*audioConnection)

type audioConnection struct {
	guild           *discordgo.Guild
	sounds          map[string]*audioClip
	soundQueue      chan string
	voiceConnection *discordgo.VoiceConnection
}

var (
	sounds           = make(map[string]*audioClip, 0)
	soundQueue       = []string{}
	soundPlayingLock = false
	voiceConnection  *discordgo.VoiceConnection
)

type audioClip struct {
	Name      string
	Extension string
	Content   [][]byte
}

// SoundsHandler -
func SoundsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// get guild ID and check for connection instance
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		// Could not find channel.
		fmt.Println("Unable to find channel.")
		return
	}

	if _, ok := activeConnections[c.GuildID]; !ok {
		newConnectionInstance, err := getNewConnectionInstance(s, m)
		if err != nil {
			log.Println(err)
			return
		}

		activeConnections[c.GuildID] = newConnectionInstance
	}

	// check if valid command
	if strings.HasPrefix(m.Content, config.Config.BotPrefix) {

		command := strings.TrimPrefix(m.Content, config.Config.BotPrefix)

		switch command {

		case "summon":
			summon(s, m)

		case "dismiss":
			dismiss()

		default:
			playAudio(command, s, m)
		}
	}
}

func getNewConnectionInstance(s *discordgo.Session, m *discordgo.MessageCreate) (*audioConnection, error) {
	return &audioConnection{}, nil
}

func dismiss() {
	if voiceConnection != nil {
		voiceConnection.Disconnect()
	}
}

func summon(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Join the channel the user issued the command from if not in it
	if voiceConnection == nil || voiceConnection.ChannelID != m.ChannelID {
		var err error

		// Find the channel that the message came from.
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			// Could not find channel.
			fmt.Println("User channel not found.")
			return
		}

		// Find the guild for that channel.
		g, err := s.State.Guild(c.GuildID)
		if err != nil {
			log.Println(err)
			return
		}

		// Look for the message sender in that guilds current voice states.
		for _, vs := range g.VoiceStates {
			if vs.UserID == m.Author.ID {

				voiceConnection, err = s.ChannelVoiceJoin(g.ID, vs.ChannelID, false, false)

				if err != nil {
					log.Println(err)
				}

				return
			}
		}

	}
}

func playAudio(soundName string, s *discordgo.Session, m *discordgo.MessageCreate) {

	// check if sound exists in memory
	if _, ok := sounds[soundName]; !ok {
		// try to load the sound if not found in memory
		err := loadFile(soundName)

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// add sound to queue
	soundQueue = append(soundQueue, soundName)

	// return if a sound is playing - it will play if it's in the queue
	if soundPlayingLock {
		return
	}

	// Find the channel that the message came from.
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		// Could not find channel.
		fmt.Println("User channel not found.")
		return
	}

	// Find the guild for that channel.
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		// Could not find guild.
		return
	}

	// Look for the message sender in that guilds current voice states.
	for _, vs := range g.VoiceStates {
		if vs.UserID == m.Author.ID {
			err = playSounds(s, g.ID, vs.ChannelID)
			if err != nil {
				fmt.Println("Error playing sound:", err)
			}

			return
		}
	}
}

// load dca file into memory
func loadFile(fileName string) error {

	// scan directory for file
	files, _ := ioutil.ReadDir(config.Config.SoundsPath)
	var fextension string
	var fname string
	for _, f := range files {
		fname = strings.Split(f.Name(), ".")[0]
		fextension = "." + strings.Split(f.Name(), ".")[1]

		if fname == fileName {
			break
		}

		fname = ""
	}

	if fname == "" {
		return errors.New("File not found")
	}

	fmt.Println("Loading file: " + fname + fextension)

	var ffmpegExecutable string

	switch runtime.GOOS {
	case "darwin":
		ffmpegExecutable = "./ffmpeg_mac"
	case "linux":
		ffmpegExecutable = "./ffmpeg_linux"
	case "windows":
		ffmpegExecutable = "ffmpeg_windows.exe"
	}

	// use ffmpeg to convert file into a format we can use
	cmd := exec.Command(ffmpegExecutable, "-i", config.Config.SoundsPath+fname+fextension, "-f", "s16le", "-ar", strconv.Itoa(frameRate), "-ac", strconv.Itoa(channels), "pipe:1")

	ffmpegout, err := cmd.StdoutPipe()

	if err != nil {
		return errors.New("Unable to execute ffmpeg. To set permissions on this file run chmod +x ffmpeg_linux (or ffmpeg_mac depending which operating system you are on)")
	}

	ffmpegbuf := bufio.NewReaderSize(ffmpegout, 16348)

	err = cmd.Start()

	if err != nil {
		return errors.New("Unable to execute ffmpeg. To set permissions on this file run chmod +x ffmpeg_linux (or ffmpeg_mac depending which operating system you are on)")
	}

	// crate encoder to convert audio to opus codec
	opusEncoder, err := gopus.NewEncoder(frameRate, channels, gopus.Audio)

	if err != nil {
		return errors.New("NewEncoder error.")
	}

	sounds[fileName] = &audioClip{
		Content:   make([][]byte, 0),
		Name:      fileName,
		Extension: fextension,
	}

	for {
		// read data from ffmpeg stdout
		audiobuf := make([]int16, frameSize*channels)
		err = binary.Read(ffmpegbuf, binary.LittleEndian, &audiobuf)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return nil
		}
		if err != nil {
			return errors.New("Error reading from ffmpeg stdout.")
		}

		// convert audio to opus codec
		opus, err := opusEncoder.Encode(audiobuf, frameSize, maxBytes)
		if err != nil {
			return errors.New("Encoding error.")
		}

		// append sound bytes to the content for this audio file
		sounds[fileName].Content = append(sounds[fileName].Content, opus)
	}

}

// playSounds - plays the current buffer to the provided channel.
func playSounds(s *discordgo.Session, guildID, channelID string) (err error) {

	//prevent other sounds from interrupting
	soundPlayingLock = true

	// Join the channel the user issued the command from if not in it
	if voiceConnection == nil || !voiceConnection.Ready {
		var err error
		voiceConnection, err = s.ChannelVoiceJoin(guildID, channelID, false, false)
		if err != nil {
			return err
		}
	}

	// keep playing sounds as long as they exist in queue
	for len(soundQueue) > 0 {

		// Sleep for a specified amount of time before playing the sound
		time.Sleep(50 * time.Millisecond)

		// Start speaking.
		_ = voiceConnection.Speaking(true)

		// Send the buffer data.
		for _, buff := range sounds[soundQueue[0]].Content {
			voiceConnection.OpusSend <- buff
		}

		// Stop speaking
		_ = voiceConnection.Speaking(false)

		// Sleep for a specificed amount of time before ending.
		time.Sleep(50 * time.Millisecond)

		soundQueue = append(soundQueue[1:])

	}

	soundPlayingLock = false

	return nil
}
