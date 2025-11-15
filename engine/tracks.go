package engine

import (
	midi "github.com/lbwise/audiowrld/mididriver"
	inst "github.com/lbwise/audiowrld/simplesynth/oscillator"
)

type TrackType int

const (
	MidiTrack TrackType = iota
	AudioTrack
)

type Track interface {
	Pipe(buffer AudioBuffer) error
	Type() TrackType
}

type InputTrack struct {
	buf       chan<- AudioBuffer
	channel   *midi.Channel
	trackType TrackType
}

func (tr *InputTrack) Type() TrackType {
	return tr.trackType
}

func (tr *InputTrack) Pipe(buffer AudioBuffer) error {
	tr.buf <- buffer
	return nil
}

func NewAudioBuffer() *AudioBuffer { return &AudioBuffer{} }

type AudioBuffer []float32

func NewAudioEngine() *Engine {
	return &Engine{
		channels:    make([]*midi.Channel, midi.MaxChannels),
		instruments: []inst.Instrument{},
		params:      NewDefaultParams(),
	}
}
