package util

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"github.com/hraban/opus"
)

// GetFileOpusData - uses ffmpeg to convert any audio
// file to opus data ready to send to discord
func GetFileOpusData(filePath string, channels, opusFrameSize, sampleRate int) ([][]byte, error) {
	cmd := exec.Command("ffmpeg", "-i", filePath, "-f", "s16le", "-ar", strconv.Itoa(sampleRate), "-ac", strconv.Itoa(channels), "pipe:1")

	cmdout, err := cmd.StdoutPipe()

	if err != nil {
		return nil, err
	}

	pcmdata := bufio.NewReader(cmdout)

	err = cmd.Start()

	if err != nil {
		return nil, err
	}

	// crate encoder to convert audio to opus codec
	opusEncoder, err := opus.NewEncoder(sampleRate, channels, opus.AppVoIP)

	if err != nil {
		return nil, errors.New("new opus encoder error")
	}

	opusOutput := make([][]byte, 0)

	for {
		// read pcm data from ffmpeg stdout
		audiobuf := make([]int16, opusFrameSize*channels)
		err = binary.Read(pcmdata, binary.LittleEndian, &audiobuf)

		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return opusOutput, nil
		}

		if err != nil {
			return nil, errors.New("error reading from ffmpeg stdout")
		}

		// convert raw pcm to opus
		opus := make([]byte, 1000)
		n, err := opusEncoder.Encode(audiobuf, opus)

		if err != nil {
			return nil, errors.New("encoding error")
		}

		// append bytes to output
		opusOutput = append(opusOutput, opus[:n])
	}
}

// GetFileExtension -
// scan directory for filename and return first extension found for that name
func GetFileExtension(path, fileName string) (string, error) {

	files, _ := ioutil.ReadDir(path)
	var fextension string
	var fname string
	for _, f := range files {
		fname = strings.Split(f.Name(), ".")[0]
		fextension = "." + strings.Split(f.Name(), ".")[1]

		if fname == fileName {
			return fextension, nil
		}
	}

	return "", errors.New("file not found")
}

// OpusToPCM - convert opus to pcm
func OpusToPCM(opusData [][]byte, sampleRate, channels int) ([]int16, error) {
	opusDecoder, err := opus.NewDecoder(sampleRate, channels)
	pcm := []int16{}

	if err != nil {
		return pcm, err
	}

	for _, o := range opusData {
		// create pcm slice with more than enough space
		p := make([]int16, 10000)
		n, err := opusDecoder.Decode(o, p)

		if err != nil {
			return pcm, err
		}

		pcm = append(pcm, p[:n*channels]...)
	}

	return pcm, nil
}

// SavePCMToWavFile - save pcm data to wav file
func SavePCMToWavFile(data []int16, filename string, sampleRate, channels int) error {

	out, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer out.Close()

	// 48 kHz, 16 bit, 2 channel, WAV.
	e := wav.NewEncoder(out, sampleRate, 16, channels, 1)

	output := new(bytes.Buffer)

	binary.Write(output, binary.LittleEndian, data)

	newReader := bytes.NewReader(output.Bytes())
	// Create new audio.IntBuffer.
	audioBuf, err := newAudioIntBuffer(newReader, sampleRate, channels)

	if err != nil {
		return err
	}

	// Write buffer to output file. This writes a RIFF header and the PCM chunks from the audio.IntBuffer.
	if err := e.Write(audioBuf); err != nil {
		return err
	}

	err = e.Close()
	return err
}

func newAudioIntBuffer(r io.Reader, sampleRate, channels int) (*audio.IntBuffer, error) {
	buf := &audio.IntBuffer{
		Format: &audio.Format{
			NumChannels: channels,
			SampleRate:  sampleRate,
		},
	}
	for {
		var sample int16
		err := binary.Read(r, binary.LittleEndian, &sample)
		switch {
		case err == io.EOF:
			return buf, nil
		case err != nil:
			return nil, err
		}
		buf.Data = append(buf.Data, int(sample))
	}
}
