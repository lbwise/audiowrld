package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/lbwise/audiowrld/instrument"
	midi "github.com/lbwise/audiowrld/mididriver"
	"github.com/lbwise/audiowrld/processing"

	"github.com/lbwise/audiowrld/io"
)

func main() {
	var chs midi.Channels = make([]*midi.Channel, midi.MaxChannels)
	err, chId := chs.NewChannel(512)
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

	stave := instrument.Stave{
		instrument.StaveNote{Note: "C", Octave: 3, Interval: 1000, Velocity: 100},
		instrument.StaveNote{Note: "D", Octave: 3, Interval: 1000, Velocity: 100},
		instrument.StaveNote{Note: "E", Octave: 3, Interval: 1000, Velocity: 100},
		instrument.StaveNote{Note: "F", Octave: 3, Interval: 1000, Velocity: 100},
		instrument.StaveNote{Note: "G", Octave: 3, Interval: 1000, Velocity: 100},
		instrument.StaveNote{Note: "A", Octave: 3, Interval: 1000, Velocity: 100},
		instrument.StaveNote{Note: "B", Octave: 3, Interval: 1000, Velocity: 100},
		instrument.StaveNote{Note: "C", Octave: 4, Interval: 1000, Velocity: 100},
	}

	// Prepares the stave to be played
	sampleSize := stave.Generate()
	fmt.Println(stave)

	buf := make([]int16, sampleSize)
	var writeIdx int
	for _, note := range stave {
		osc := instrument.SinOscillator{Note: note}
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
	//io.WavPlot(fileName, buf)
	io.CreateWAV("./exports", fileName, buf, sampleSize)
}
