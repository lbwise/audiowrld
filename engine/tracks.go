package engine

import (
	"github.com/lbwise/audiowrld/audio"
	midi "github.com/lbwise/audiowrld/mididriver"
	inst "github.com/lbwise/audiowrld/simplesynth/oscillator"
)

type TrackType int

const (
	MidiTrack TrackType = iota
	AudioTrack
)

type Track interface {
	Pipe(buffer audio.Buffer) error
	Type() TrackType
}

type InputTrack struct {
	buf       chan<- audio.Buffer
	channel   *midi.Channel
	trackType TrackType
}

func (tr *InputTrack) Type() TrackType {
	return tr.trackType
}

func (tr *InputTrack) Pipe(buffer audio.Buffer) error {
	tr.buf <- buffer
	return nil
}

func NewAudioEngine() *Engine {
	return &Engine{
		channels:    make([]*midi.Channel, midi.MaxChannels),
		instruments: []inst.Instrument{},
		params:      audio.NewDefaultParams(),
	}
}
