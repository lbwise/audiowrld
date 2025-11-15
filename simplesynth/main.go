package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"time"

	midi "github.com/lbwise/audiowrld/mididriver"
	"github.com/lbwise/audiowrld/processing"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/lbwise/audiowrld/simplesynth/io"
	"github.com/lbwise/audiowrld/simplesynth/notes"
	"github.com/lbwise/audiowrld/simplesynth/oscillator"
)

func main() {
	var chs midi.Channels = make([]*midi.Channel, midi.MaxChannels)
	err, chId := chs.NewChannel()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var stopCb func() error
	err, stopCb = chs[chId].Play(512)
	if err != nil {
		panic(err)
	}

	in := make(chan midi.RawMsg, 64)
	scn := midi.NewScanner(in, chs)
	go func() {
		defer wg.Done()
		err := scn.Scan()
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		defer wg.Done()
		defer func() {
			stopCb()
		}()
		midi.MidiSim(in, int(chId))
	}()

	if err != nil {
		panic(err)
	}

	wg.Wait()
	return

	if len(os.Args) > 1 && os.Args[1] == "midi" {
		return
	}

	stave := notes.Stave{
		notes.StaveNote{Note: "C", Octave: 3, Interval: 1000, Velocity: 100},
		notes.StaveNote{Note: "D", Octave: 3, Interval: 1000, Velocity: 100},
		notes.StaveNote{Note: "E", Octave: 3, Interval: 1000, Velocity: 100},
		notes.StaveNote{Note: "F", Octave: 3, Interval: 1000, Velocity: 100},
		notes.StaveNote{Note: "G", Octave: 3, Interval: 1000, Velocity: 100},
		notes.StaveNote{Note: "A", Octave: 3, Interval: 1000, Velocity: 100},
		notes.StaveNote{Note: "B", Octave: 3, Interval: 1000, Velocity: 100},
		notes.StaveNote{Note: "C", Octave: 4, Interval: 1000, Velocity: 100},
	}

	// Prepares the stave to be played
	sampleSize := stave.Generate()
	fmt.Println(stave)

	buf := make([]int16, sampleSize)
	var writeIdx int
	for _, note := range stave {
		osc := oscillator.SinOscillator{Note: note}
		n, err := osc.Generate(buf, writeIdx)
		writeIdx += n
		if err != nil {
			panic("COULD NOT GENERATE SOUND")
		}
	}

	ce := processing.NewClippingEffect(120, 40)
	ce.Bypass = true
	//err := ce.Process(buf, buf)
	if err != nil {
		panic(err)
	}

	fmt.Println(buf[:20])

	soundTitle := "test-sin-major"
	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("%s-%d", soundTitle, timestamp)
	//wavPlot(fileName, buf)
	io.CreateWAV("./exports", fileName, buf, sampleSize)
}

func wavPlot(title string, amplitudes []int16) {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "Samp. no"
	p.Y.Label.Text = "Volume"

	var max int16
	for _, amplitude := range amplitudes {
		abs := int16(math.Abs(float64(amplitude)))
		if abs > max {
			max = abs
		}
	}

	const sampleFactor = 300
	dsN := len(amplitudes) / sampleFactor
	xs := make([]float64, dsN)
	ys := make([]float64, dsN)

	for i := range amplitudes {
		dsIdx := i / sampleFactor

		if dsIdx == 0 {
			xs[dsIdx] = float64(dsIdx)

			if max != 0 {
				ys[dsIdx] = float64(amplitudes[i]) / float64(max)
			} else {
				ys[dsIdx] = float64(amplitudes[i])
			}
		}
	}

	//err := plotutil.AddScatters(p, hplot.ZipXY(xs, ys))
	//if err != nil {
	//	fmt.Println(fmt.Errorf("could not create scatters: %+v", err))
	//}

	pts := make(plotter.XYs, dsN*2)
	for i := 0; i < len(amplitudes); i += sampleFactor {
		x := float64(i / sampleFactor)
		y := float64(amplitudes[i]) / float64(max)
		pts[(i/sampleFactor)*2] = plotter.XY{X: x, Y: 0}
		pts[(i/sampleFactor)*2+1] = plotter.XY{X: x, Y: y}
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		fmt.Println(fmt.Errorf("could not create line plot: %+v", err))
	}

	p.Add(line)

	err = p.Save(20*vg.Centimeter, 10*vg.Centimeter, fmt.Sprintf("exports/%s-waveform.png", title))
	if err != nil {
		fmt.Println(fmt.Errorf("could not save scatter plot: %+v", err))
	}
}
