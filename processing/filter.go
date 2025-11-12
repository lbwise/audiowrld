package processing

import (
	"github.com/lbwise/audiowrld/simplesynth/constants"
)

//func ApplyLPFilter(data bin, cutoffRate int) {
//
//}

type Effect interface {
	Process(input, output []int16) error
}

func NewClippingEffect(gain, level uint8) *ClippingEffect {
	return &ClippingEffect{Bypass: false, Gain: gain, Level: level, gainMin: 10, gainMax: 110}
}

type ClippingAlgorithm int

const (
	Distortion ClippingAlgorithm = iota
	Overdrive
	Saturation
	Fuzz
)

// Fix gain min and max so that the gain property adjusts
// the propertion between the two instead of a cutoff.
type ClippingEffect struct {
	Bypass  bool
	Gain    uint8
	Level   uint8
	Algo    ClippingAlgorithm
	gainMin uint8
	gainMax uint8
}

func (ce *ClippingEffect) Process(in, out []int16) []int16 {
	if ce.Bypass {
		copy(out, in)
		return nil
	}
	gain := min(max(ce.Gain, ce.gainMin), ce.gainMax)
	clipVal := (1 - (float64(gain) / 128.0)) * float64(constants.MaxAmp)
	level := 1.0 + (float64(ce.Level) / 128.0)

	for i := 0; i < len(in); i += 1 {
		amp := float64(in[i])
		if amp > clipVal {
			amp = clipVal
		} else if amp < -clipVal {
			amp = -clipVal
		}

		amp *= level
		maxAmp := float64(constants.MaxAmp)
		if amp > maxAmp {
			amp = maxAmp
		} else if amp < -maxAmp {
			amp = -maxAmp
		}
		out[i] = int16(amp)
	}
	return nil
}

//func ExtractFreqs(buf []int16) []int {
//
//}
//
//func normalizeBuffer(in, out []int16) error {
//	maxAmp := 0.0
//	for i := range in {
//		if in[i] > maxAmp {
//			maxAmp = in[i]
//		}
//
//	}
//}
