package handlers

import (
	"../config"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"os"
	"strings"
	"time"
)

var (
	sounds = make(map[string][][]byte, 0)

	soundPlayingLock = false
)

const SOUNDS_DIR string = "./sounds/"

func SoundsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// exit function call if sound is playing
	if soundPlayingLock {
		fmt.Println("Exiting function call")
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
	fmt.Println("Loading file: " + fileName + ".dca")

	file, err := os.Open(SOUNDS_DIR + fileName + ".dca")

	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return err
	}

	sounds[fileName] = make([][]byte, 0)

	var opuslen int16

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err != nil {
			file.Close()
			if err == io.EOF {
				return nil
			} else if err == io.ErrUnexpectedEOF {
				return err
			}
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

		sounds[fileName] = append(sounds[fileName], InBuf)
	}

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

	soundPlayingLock = false

	return nil
}
