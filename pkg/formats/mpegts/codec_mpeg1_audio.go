package mpegts

import (
	"github.com/asticode/go-astits"
)

// CodecMPEG1Audio is a MPEG-1 Audio codec.
type CodecMPEG1Audio struct{}

func (c CodecMPEG1Audio) marshal(pid uint16) (*astits.PMTElementaryStream, error) {
	return &astits.PMTElementaryStream{
		ElementaryPID:               pid,
		ElementaryStreamDescriptors: nil,
		StreamType:                  astits.StreamTypeMPEG1Audio,
	}, nil
}

// IsVideo implements Codec.
func (CodecMPEG1Audio) IsVideo() bool {
	return true
}
