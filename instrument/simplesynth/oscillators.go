package instrument

import (
	"math"

	"github.com/lbwise/audiowrld/audio"
	"github.com/lbwise/audiowrld/instrument"
)

type Oscillator interface {
	Generate([]int16, int) (int, error)
}

type SquareOscillator struct {
	Note instrument.StaveNote
}

func (s *SquareOscillator) Generate(buf []int16, writeIdx int) (int, error) {

	numSamples := s.Note.Interval * audio.SampleRate / 1000
	amp := s.Note.Amplitude

	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(audio.SampleRate)
		y := float64(amp) * math.Sin(2*math.Pi*s.Note.Frequency*t)
		if y >= 0 {
			buf[writeIdx+i] = amp
		} else {
			buf[writeIdx+i] = -amp
		}
	}
	return numSamples, nil
}

type SinOscillator struct {
	Note instrument.StaveNote
}

func (s *SinOscillator) Generate(buf []int16, writeIdx int) (int, error) {
	numSamples := s.Note.Interval * audio.SampleRate / 1000
	amp := s.Note.Amplitude

	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(audio.SampleRate)
		y := float64(amp) * math.Sin(2*math.Pi*s.Note.Frequency*t)
		buf[writeIdx+i] = int16(y)
	}
	return numSamples, nil
}

type TriangleOscillator struct {
	Note instrument.StaveNote
}

func (s *TriangleOscillator) Generate(buf []int16, writeIdx int) (int, error) {
	numSamples := s.Note.Interval * audio.SampleRate / 1000
	amp := float64(s.Note.Amplitude)

	var phase float64
	// y = 4/f * x - amp
	for i := 0; i < numSamples; i++ {
		// I NEED THIS EXPLAINED
		phase += s.Note.Frequency / float64(audio.SampleRate)
		if phase >= 1.0 {
			phase -= 1.0
		}

		// triangle wave formula
		y := 4*amp*math.Abs(phase-0.5) - amp
		buf[writeIdx+i] = int16(y)
	}
	return numSamples, nil
}
