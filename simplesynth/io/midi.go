package io

import "github.com/lbwise/audiowrld/simplesynth/notes"

type MidiInput struct {
	NumSamples int
	Notes      []MidiNote
	BendVal    int
}

type MidiNote struct {
	Note     notes.StaveNote
	Octave   int
	Velocity int
}

func NewMidiNote(note notes.StaveNote, octave int, velocity int) *MidiNote {
	return &MidiNote{note, octave, velocity}
}
