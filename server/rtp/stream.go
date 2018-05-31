package rtp

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// Stream -
type Stream struct {
	StartTime         time.Time
	StartingTimestamp uint32
	LatestTimestamp   uint32
	Packets           []*discordgo.Packet
}

// GetTimeDiff -
func (s *Stream) GetTimeDiff() uint32 {
	return s.LatestTimestamp - s.StartingTimestamp
}

func (s *Stream) GetTimeOffset(startTime int64) int64 {
	return s.StartTime.Unix() - startTime
}

// PushPacket -
func (s *Stream) PushPacket(packet *discordgo.Packet) {

	if s.StartTime.Unix() == 0 {
		s.StartTime = time.Now()
	}

	if s.StartingTimestamp == 0 && packet.Timestamp != 0 {
		s.StartingTimestamp = packet.Timestamp
	}

	s.LatestTimestamp = packet.Timestamp

	s.Packets = append(s.Packets, packet)
}
