package instrument

type MidiInput struct {
	NumSamples int
	Notes      []MidiNote
	BendVal    int
}

type MidiNote struct {
	Note     StaveNote
	Octave   int
	Velocity int
}
