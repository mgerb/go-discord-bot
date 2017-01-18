package handlers

import (
	"../config"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var sounds = make(map[string][][]byte, 0)

const SOUNDS_DIR string = "./sounds/"

func SoundsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.Config.Activator) {
		// Find the channel that the message came from.
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			// Could not find channel.
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
				err = playSound(s, g.ID, vs.ChannelID, strings.TrimPrefix(m.Content, config.Config.Activator))
				if err != nil {
					fmt.Println("Error playing sound:", err)
				}

				return
			}
		}
	}
}

// loadSound attempts to load an encoded sound file from disk.
func LoadSounds() error {

	files, _ := ioutil.ReadDir(SOUNDS_DIR)

	for _, file := range files {
		fmt.Println(file.Name())
		err := loadFile(file.Name())

		if err != nil {
			fmt.Println(err.Error())
		}

	}

	return nil
}

func loadFile(fileName string) error {

	trimmedName := strings.TrimSuffix(fileName, ".dca")
	sounds[trimmedName] = make([][]byte, 0)

	file, err := os.Open(SOUNDS_DIR + fileName)

	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return err
	}

	var opuslen int16

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			file.Close()
			if err != nil {
				return err
			}
			return nil
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		sounds[trimmedName] = append(sounds[trimmedName], InBuf)
	}

}

// playSound plays the current buffer to the provided channel.
func playSound(s *discordgo.Session, guildID, channelID string, sound string) (err error) {

	if _, ok := sounds[sound]; !ok {
		return errors.New("Sound not found")
	}

	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, false)
	if err != nil {
		return err
	}

	// Sleep for a specified amount of time before playing the sound
	time.Sleep(250 * time.Millisecond)

	// Start speaking.
	_ = vc.Speaking(true)

	// Send the buffer data.
	for _, buff := range sounds[sound] {
		vc.OpusSend <- buff
	}

	// Stop speaking
	_ = vc.Speaking(false)

	// Sleep for a specificed amount of time before ending.
	time.Sleep(250 * time.Millisecond)

	// Disconnect from the provided voice channel.
	_ = vc.Disconnect()

	return nil
}
