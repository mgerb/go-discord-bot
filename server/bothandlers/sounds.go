package bothandlers

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"layeh.com/gopus"

	"github.com/bwmarrin/discordgo"
	"github.com/mgerb/go-discord-bot/server/config"
	log "github.com/sirupsen/logrus"
)

const (
	channels                 int = 2                   // 1 for mono, 2 for stereo
	sampleRate               int = 48000               // audio sampling rate - apparently a standard for opus
	frameSize                int = 960                 // uint16 size of each audio frame
	maxBytes                 int = (frameSize * 2) * 2 // max size of opus data
	maxSoundQueue            int = 10                  // max amount of sounds that can be queued at one time
	voiceClipQueuePacketSize int = 2000                // this packet size equates to roughly 40 seconds of audio
)

// ActiveConnections - current active bot connections
// store our connection objects in a map tied to a guild id
var ActiveConnections = make(map[string]*AudioConnection)
var speakers = make(map[uint32]*gopus.Decoder)

// AudioConnection -
type AudioConnection struct {
	Guild             *discordgo.Guild           `json:"guild"`
	Session           *discordgo.Session         `json:"-"`
	VoiceConnection   *discordgo.VoiceConnection `json:"-"`
	CurrentChannel    *discordgo.Channel         `json:"current_channel"`
	Sounds            map[string]*AudioClip      `json:"-"`
	SoundQueue        chan string                `json:"-"`
	VoiceClipQueue    chan *discordgo.Packet     `json:"-"`
	SoundPlayingLock  bool                       `json:"-"`
	AudioListenerLock bool                       `json:"-"`
	Mutex             *sync.Mutex                `json:"-"` // mutex for single audio connection
}

type AudioClip struct {
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
		log.Error("Unable to find channel.")
		return
	}

	// check to see if active connection object exists
	if _, ok := ActiveConnections[c.GuildID]; !ok {

		// Find the guild for that channel.
		newGuild, err := s.State.Guild(c.GuildID)

		if err != nil {
			log.Error(err)
			return
		}

		// create new connection instance
		ActiveConnections[c.GuildID] = &AudioConnection{
			Guild:             newGuild,
			Session:           s,
			Sounds:            make(map[string]*AudioClip, 0),
			SoundQueue:        make(chan string, maxSoundQueue),
			Mutex:             &sync.Mutex{},
			AudioListenerLock: false,
		}
	}

	ActiveConnections[c.GuildID].handleMessage(m)
}

func (conn *AudioConnection) handleMessage(m *discordgo.MessageCreate) {

	// check if valid command
	if strings.HasPrefix(m.Content, config.Config.BotPrefix) {

		command := strings.TrimPrefix(m.Content, config.Config.BotPrefix)

		switch command {

		case "summon":
			conn.summon(m)

		case "dismiss":
			conn.dismiss()

		case "clip":
			conn.clipAudio(m)

		case "random":
			conn.playRandomAudio(m)

		default:
			conn.PlayAudio(command, m)
		}
	}
}

// dismiss bot from currnet channel if it's in one
func (conn *AudioConnection) dismiss() {
	if conn.VoiceConnection != nil && !conn.SoundPlayingLock && len(conn.SoundQueue) == 0 {
		conn.VoiceConnection.Disconnect()
	}
}

// summon bot to channel that user is currently in
func (conn *AudioConnection) summon(m *discordgo.MessageCreate) {

	// Join the channel the user issued the command from if not in it
	if conn.VoiceConnection == nil || conn.VoiceConnection.ChannelID != m.ChannelID {

		var err error

		// Find the channel that the message came from.
		c, err := conn.Session.State.Channel(m.ChannelID)
		if err != nil {
			// Could not find channel.
			log.Error("User channel not found.")
			return
		}

		// Find the guild for that channel.
		g, err := conn.Session.State.Guild(c.GuildID)
		if err != nil {
			log.Error(err)
			return
		}

		// Look for the message sender in that guilds current voice states
		for _, vs := range g.VoiceStates {
			if vs.UserID == m.Author.ID {

				conn.VoiceConnection, err = conn.Session.ChannelVoiceJoin(g.ID, vs.ChannelID, false, false)

				if err != nil {
					log.Error(err)
					return
				}

				// set the current channel
				conn.CurrentChannel = c

				// start listening to audio if not locked
				if !conn.AudioListenerLock {
					go conn.startAudioListener()
				}
			}
		}

	}
}

func (conn *AudioConnection) queueAudio(soundName string) {
}

// play a random sound clip
func (conn *AudioConnection) playRandomAudio(m *discordgo.MessageCreate) {
	files, _ := ioutil.ReadDir(config.Config.SoundsPath)
	if len(files) > 0 {
		randomIndex := rand.Intn(len(files))
		arr := strings.Split(files[randomIndex].Name(), ".")
		if len(arr) > 0 && arr[0] != "" {
			conn.PlayAudio(arr[0], m)
		}
	}
}

// PlayAudio - play audio in channel that user is in
// if MessageCreate is null play in current channel
func (conn *AudioConnection) PlayAudio(soundName string, m *discordgo.MessageCreate) {

	// summon bot to channel if new message passed in
	if m != nil {
		conn.summon(m)
	} else if !conn.VoiceConnection.Ready {
		return
	}

	// check if sound exists in memory
	if _, ok := conn.Sounds[soundName]; !ok {
		// try to load the sound if not found in memory
		err := conn.loadFile(soundName)

		if err != nil {
			log.Error(err)
			return
		}
	}

	// add sound to queue if queue isn't full
	select {
	case conn.SoundQueue <- soundName:

	default:
		break
	}

	// start playing sounds in queue if not already playing
	if !conn.SoundPlayingLock {
		conn.playSoundsInQueue()
	}

}

// playSoundsInQueue - play sounds until audio queue is empty
func (conn *AudioConnection) playSoundsInQueue() {
	conn.toggleSoundPlayingLock(true)

	// Start speaking.
	_ = conn.VoiceConnection.Speaking(true)

	for {
		select {
		case newSoundName := <-conn.SoundQueue:

			if !conn.VoiceConnection.Ready {
				return
			}

			// Send the buffer data.
			for _, buff := range conn.Sounds[newSoundName].Content {
				conn.VoiceConnection.OpusSend <- buff
			}

			// Sleep for a specificed amount of time before ending.
			time.Sleep(100 * time.Millisecond)

		default:
			// Stop speaking
			_ = conn.VoiceConnection.Speaking(false)
			conn.toggleSoundPlayingLock(false)
			return
		}
	}
}

func (conn *AudioConnection) toggleSoundPlayingLock(playing bool) {
	conn.Mutex.Lock()
	conn.SoundPlayingLock = playing
	conn.Mutex.Unlock()
}

// load audio file into memory
func (conn *AudioConnection) loadFile(fileName string) error {

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

	log.Debug("Loading file: " + fname + fextension)

	// use ffmpeg to convert file into a format we can use
	cmd := exec.Command("ffmpeg", "-i", config.Config.SoundsPath+"/"+fname+fextension, "-f", "s16le", "-ar", strconv.Itoa(sampleRate), "-ac", strconv.Itoa(channels), "pipe:1")

	ffmpegout, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	ffmpegbuf := bufio.NewReaderSize(ffmpegout, 16348)

	err = cmd.Start()

	if err != nil {
		return err
	}

	// crate encoder to convert audio to opus codec
	opusEncoder, err := gopus.NewEncoder(sampleRate, channels, gopus.Audio)

	if err != nil {
		return errors.New("NewEncoder error.")
	}

	conn.Sounds[fileName] = &AudioClip{
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
		conn.Sounds[fileName].Content = append(conn.Sounds[fileName].Content, opus)
	}

}

func (conn *AudioConnection) clipAudio(m *discordgo.MessageCreate) {
	if len(conn.VoiceClipQueue) < 10 {
		conn.Session.ChannelMessageSend(m.ChannelID, "Clip failed.")
	} else {
		writePacketsToFile(m.Author.Username, conn.VoiceClipQueue)
		conn.Session.ChannelMessageSend(m.ChannelID, "Sound clipped!")
	}
}

func writePacketsToFile(username string, packets chan *discordgo.Packet) {

	// create clips folder if it does not exist
	if _, err := os.Stat(config.Config.ClipsPath); os.IsNotExist(err) {
		os.Mkdir(config.Config.ClipsPath, os.ModePerm)
	}

	// construct filename
	timestamp := time.Now().UTC().Format("2006-01-02") + "-" + strconv.Itoa(int(time.Now().Unix()))
	filename := config.Config.ClipsPath + "/" + timestamp + "-" + username + ".wav"

	// grab everything from the voice packet channel and dump it to the file
	// close when there is nothing left
	pcmOut := make([]int16, 0)

loop:
	for {
		select {
		case p := <-packets:
			for _, pcm := range p.PCM {
				pcmOut = append(pcmOut, pcm)
			}
		default:
			break loop
		}
	}

	cmd := exec.Command("ffmpeg", "-f", "s16le", "-ar", strconv.Itoa(sampleRate), "-ac", strconv.Itoa(channels), "-i", "pipe:0", filename)

	output := new(bytes.Buffer)

	binary.Write(output, binary.LittleEndian, pcmOut)
	cmd.Stdin = bytes.NewReader(output.Bytes())

	err := cmd.Run()

	if err != nil {
		log.Error(err)
	}
}

// start listening to the voice channel
func (conn *AudioConnection) startAudioListener() {

	conn.AudioListenerLock = true

	if conn.VoiceClipQueue == nil {
		conn.VoiceClipQueue = make(chan *discordgo.Packet, voiceClipQueuePacketSize)
	}

loop:
	for {

		select {
		// grab incomming audio
		case opusChannel, ok := <-conn.VoiceConnection.OpusRecv:
			if !ok {
				continue
			}

			var err error
			_, ok = speakers[opusChannel.SSRC]

			if !ok {
				speakers[opusChannel.SSRC], err = gopus.NewDecoder(sampleRate, channels)
				if err != nil {
					log.Error("error creating opus decoder", err)
					continue
				}
			}

			opusChannel.PCM, err = speakers[opusChannel.SSRC].Decode(opusChannel.Opus, frameSize, false)
			if err != nil {
				log.Error("Error decoding opus data", err)
				continue
			}

			// if channel is full trim off from beginning
			if len(conn.VoiceClipQueue) == cap(conn.VoiceClipQueue) {
				<-conn.VoiceClipQueue
			}

			// add current packet to channel queue
			conn.VoiceClipQueue <- opusChannel

		// check if voice connection fails then break out of audio listener
		default:
			if !conn.VoiceConnection.Ready {
				break loop
			}
			time.Sleep(time.Second * 1)
		}
	}

	// remove lock upon exit
	conn.AudioListenerLock = false
}
