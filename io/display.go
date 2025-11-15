package io

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func WavPlot(title string, amplitudes []int16) {
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
