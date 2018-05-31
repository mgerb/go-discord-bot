// rtp packet processing

package rtp

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// RTP -
type RTP struct {
	StartTime time.Time
	Streams   map[uint32]*Stream
}

// New -
func New() *RTP {
	return &RTP{
		StartTime: time.Now(),
		Streams:   map[uint32]*Stream{},
	}
}

// PushPacket -
func (r *RTP) PushPacket(packet *discordgo.Packet) {
	if _, ok := r.Streams[packet.SSRC]; !ok {
		r.Streams[packet.SSRC] = new(Stream)
	}
	r.Streams[packet.SSRC].PushPacket(packet)
}

// GetStream -
func (r *RTP) GetStream(ssrc uint32) *Stream {
	return r.Streams[ssrc]
}
