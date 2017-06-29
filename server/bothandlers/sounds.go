package bothandlers

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"../config"
	"github.com/bwmarrin/discordgo"
	"layeh.com/gopus"
)

const (
	channels  int = 2                   // 1 for mono, 2 for stereo
	frameRate int = 48000               // audio sampling rate
	frameSize int = 960                 // uint16 size of each audio frame
	maxBytes  int = (frameSize * 2) * 2 // max size of opus data
)

var (
	sounds           = make(map[string]*AudioClip, 0)
	soundPlayingLock = false
)

type AudioClip struct {
	Name      string
	Extension string
	Content   [][]byte
}

const SOUNDS_DIR string = "./sounds/"

func SoundsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// exit function call if sound is playing
	if soundPlayingLock {
		fmt.Println("Function in progress, exiting function call...")
		return
	}

	// check if valid command
	if strings.HasPrefix(m.Content, config.Config.BotPrefix) {

		soundName := strings.TrimPrefix(m.Content, config.Config.BotPrefix)

		// check if sound exists in memory
		if _, ok := sounds[soundName]; !ok {
			// try to load the sound if not found in memory
			err := loadFile(soundName)

			if err != nil {
				fmt.Println(err)
				return
			}
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
				err = playSound(s, g.ID, vs.ChannelID, soundName)
				if err != nil {
					fmt.Println("Error playing sound:", err)
				}

				return
			}
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

	sounds[fileName] = &AudioClip{
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

	return nil
}

// playSound plays the current buffer to the provided channel.
func playSound(s *discordgo.Session, guildID, channelID string, sound string) (err error) {

	if _, ok := sounds[sound]; !ok {
		return errors.New("Sound not found")
	}

	//prevent other sounds from interrupting
	soundPlayingLock = true

	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, false)
	if err != nil {
		return err
	}

	// Sleep for a specified amount of time before playing the sound
	time.Sleep(100 * time.Millisecond)

	// Start speaking.
	_ = vc.Speaking(true)

	// Send the buffer data.
	for _, buff := range sounds[sound].Content {
		vc.OpusSend <- buff
	}

	// Stop speaking
	_ = vc.Speaking(false)

	// Sleep for a specificed amount of time before ending.
	time.Sleep(250 * time.Millisecond)

	// Disconnect from the provided voice channel.
	_ = vc.Disconnect()

	soundPlayingLock = false

	return nil
}
