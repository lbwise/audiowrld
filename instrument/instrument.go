package instrument

type Instrument interface {
	Play(input MidiInput) []int16
	Configure(settings InstrumentSettings)
}

type InstrumentType int

const (
	SinOscillatorVal InstrumentType = iota
	SquareOscillatorVal
	TriangleOscillatorVal
)

type ParamVal int

func (p ParamVal) isValid() bool {
	return p >= 0 && p < 128
}

type InstrumentSettings struct {
	Type      InstrumentType
	Volume    ParamVal
	FilterVal ParamVal
}
