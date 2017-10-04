package bothandlers

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"layeh.com/gopus"

	"github.com/bwmarrin/discordgo"
	"github.com/cryptix/wav"
	"github.com/mgerb/go-discord-bot/server/config"
)

const (
	channels      int = 2                   // 1 for mono, 2 for stereo
	frameRate     int = 48000               // audio sampling rate
	frameSize     int = 960                 // uint16 size of each audio frame
	maxBytes      int = (frameSize * 2) * 2 // max size of opus data
	maxSoundQueue int = 10                  // max amount of sounds that can be queued at one time

	// TODO: figure out why api isn't returning bitrate
	bitRate    int = 64000
	sampleRate int = 96000
)

// store our connection objects in a map tied to a guild id
var activeConnections = make(map[string]*audioConnection)

type audioConnection struct {
	guild            *discordgo.Guild
	session          *discordgo.Session
	sounds           map[string]*audioClip
	soundQueue       chan string
	voiceConnection  *discordgo.VoiceConnection
	currentChannel   *discordgo.Channel
	soundPlayingLock bool
	mutex            *sync.Mutex // mutex for single audio connection
}

type audioClip struct {
	Name      string
	Extension string
	Content   [][]byte
}

// SoundsHandler -
func SoundsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// get channel state to get guild id
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		// Could not find channel.
		fmt.Println("Unable to find channel.")
		return
	}

	// check to see if active connection object exists
	if _, ok := activeConnections[c.GuildID]; !ok {

		// Find the guild for that channel.
		newGuild, err := s.State.Guild(c.GuildID)
		if err != nil {
			log.Println(err)
			return
		}

		// create new connection instance
		newInstance := &audioConnection{
			guild:      newGuild,
			session:    s,
			sounds:     make(map[string]*audioClip, 0),
			soundQueue: make(chan string, maxSoundQueue),
			mutex:      &sync.Mutex{},
		}

		activeConnections[c.GuildID] = newInstance

		// start listening on the sound channel
		go activeConnections[c.GuildID].playSounds()
	}

	// start new go routine handling the message
	go activeConnections[c.GuildID].handleMessage(m)
}

func (conn *audioConnection) handleMessage(m *discordgo.MessageCreate) {

	// check if valid command
	if strings.HasPrefix(m.Content, config.Config.BotPrefix) {

		command := strings.TrimPrefix(m.Content, config.Config.BotPrefix)

		switch command {

		case "summon":
			conn.summon(m)

		case "dismiss":
			conn.dismiss()

		case "clip":
			conn.clipAudio()

		default:
			conn.playAudio(command, m)
		}
	}
}

func (conn *audioConnection) dismiss() {
	if conn.voiceConnection != nil && !conn.soundPlayingLock && len(conn.soundQueue) == 0 {
		conn.voiceConnection.Disconnect()
	}
}

func (conn *audioConnection) summon(m *discordgo.MessageCreate) {
	// Join the channel the user issued the command from if not in it
	if conn.voiceConnection == nil || conn.voiceConnection.ChannelID != m.ChannelID {
		var err error

		// Find the channel that the message came from.
		c, err := conn.session.State.Channel(m.ChannelID)
		if err != nil {
			// Could not find channel.
			fmt.Println("User channel not found.")
			return
		}

		// Find the guild for that channel.
		g, err := conn.session.State.Guild(c.GuildID)
		if err != nil {
			log.Println(err)
			return
		}

		// Look for the message sender in that guilds current voice states.
		for _, vs := range g.VoiceStates {
			if vs.UserID == m.Author.ID {

				conn.voiceConnection, err = conn.session.ChannelVoiceJoin(g.ID, vs.ChannelID, false, false)

				if err != nil {
					log.Println(err)
				}

				// set the current channel
				conn.currentChannel = c

				// start listening to audio after joining channel
				go conn.startAudioListener()

				return
			}
		}

	}
}

func (conn *audioConnection) playAudio(soundName string, m *discordgo.MessageCreate) {

	// check if sound exists in memory
	if _, ok := conn.sounds[soundName]; !ok {
		// try to load the sound if not found in memory
		err := conn.loadFile(soundName)

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// summon bot to channel
	conn.summon(m)

	// add sound to queue if queue isn't full
	select {
	case conn.soundQueue <- soundName:

	default:
		return
	}

}

// load dca file into memory
func (conn *audioConnection) loadFile(fileName string) error {

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

	// use ffmpeg to convert file into a format we can use
	cmd := exec.Command("ffmpeg", "-i", config.Config.SoundsPath+fname+fextension, "-f", "s16le", "-ar", strconv.Itoa(frameRate), "-ac", strconv.Itoa(channels), "pipe:1")

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

	conn.sounds[fileName] = &audioClip{
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
		conn.sounds[fileName].Content = append(conn.sounds[fileName].Content, opus)
	}

}

func (conn *audioConnection) clipAudio() {

}

func (conn *audioConnection) startAudioListener() {

	speakers := make(map[uint32]*gopus.Decoder)
	voicePackets := []*discordgo.Packet{}
	var err error

loop:
	for {

		select {
		case opusChannel, ok := <-conn.voiceConnection.OpusRecv:
			if !ok {
				continue
			}

			_, ok = speakers[opusChannel.SSRC]

			if !ok {
				speakers[opusChannel.SSRC], err = gopus.NewDecoder(frameRate, 1)
				if err != nil {
					log.Println("error creating opus decoder", err)
					continue
				}
			}

			opusChannel.PCM, err = speakers[opusChannel.SSRC].Decode(opusChannel.Opus, frameSize, false)
			if err != nil {
				log.Println("Error decoding opus data", err)
				continue
			}

			voicePackets = append(voicePackets, opusChannel)

			// TODO:
			fmt.Println(len(voicePackets))
			fmt.Println(int(opusChannel.Timestamp) / bitRate)
			if len(voicePackets) > bitRate*5 {
				voicePackets = voicePackets[1:]
			}

		default:
			if !conn.voiceConnection.Ready {
				break loop
			}

			// cleanOldPackets(voicePackets)
		}

	}

	writeWavFile(voicePackets)

}

// // remove packets that are more than 1 minute old
// func cleanOldPackets(packets []*discordgo.Packet) {

// 	if len(packets) < 1 {
// 		return
// 	}

// 	// latest timestamp divided by bitrate
// 	currentTime := int(packets[len(packets)-1].Timestamp) / bitRate

// 	index := 0
// 	// grab the index of the first object less than a minute ago
// 	for i, p := range packets {
// 		timestamp := int(p.Timestamp) / bitRate
// 		if timestamp < currentTime-5 {
// 			index = i
// 			break
// 		} else if timestamp > currentTime-5 {
// 			break
// 		}
// 	}

// 	if index != 0 {

// 		log.Println("cleaning")
// 		log.Println(index)
// 	}

// 	packets = packets[index:]

// }

func writeWavFile(packets []*discordgo.Packet) {

	wavOut, err := os.Create("test.wav")
	checkErr(err)
	defer wavOut.Close()

	meta := wav.File{
		Channels:        1,
		SampleRate:      uint32(sampleRate),
		SignificantBits: 16,
	}

	writer, err := meta.NewWriter(wavOut)
	checkErr(err)
	defer writer.Close()

	for _, p := range packets {
		for _, pcm := range p.PCM {
			err := writer.WriteInt32(int32(pcm))
			checkErr(err)
		}
	}
}

// playSounds - plays the current buffer to the provided channel.
func (conn *audioConnection) playSounds() (err error) {

	for {
		newSoundName := <-conn.soundQueue

		conn.toggleSoundPlayingLock(true)

		if !conn.voiceConnection.Ready {
			continue
		}

		// Start speaking.
		_ = conn.voiceConnection.Speaking(true)

		// Send the buffer data.
		for _, buff := range conn.sounds[newSoundName].Content {
			conn.voiceConnection.OpusSend <- buff
		}

		// Stop speaking
		_ = conn.voiceConnection.Speaking(false)

		// Sleep for a specificed amount of time before ending.
		time.Sleep(50 * time.Millisecond)

		conn.toggleSoundPlayingLock(false)
	}

}

func (conn *audioConnection) toggleSoundPlayingLock(playing bool) {
	conn.mutex.Lock()
	conn.soundPlayingLock = playing
	conn.mutex.Unlock()
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
