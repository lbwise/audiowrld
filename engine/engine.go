package engine

import (
	"fmt"
	"sync"

	"github.com/lbwise/audiowrld/io"
	midi "github.com/lbwise/audiowrld/mididriver"
	"github.com/lbwise/audiowrld/simplesynth/constants"
	inst "github.com/lbwise/audiowrld/simplesynth/oscillator"
)

type Engine struct {
	channels     midi.Channels
	instruments  []inst.Instrument
	params       *Params
	outputDevice *io.OutputDevice
	scanChan     chan midi.RawMsg
	tracks       []Track
	master       Track
	tick         int
}

type Track interface {
	Name() string
}

type InputTrack struct {
	buf AudioBuffer
}

type AudioBuffer []float32

func NewAudioEngine() *Engine {
	return &Engine{
		channels:    make([]*midi.Channel, midi.MaxChannels),
		instruments: []inst.Instrument{},
		params:      NewDefaultParams(),
	}
}

func (eng *Engine) Init() error {
	// detect midi devices
	// start scanner

	eng.scanChan = make(chan midi.RawMsg, 64)
	scn := midi.NewScanner(eng.scanChan, eng.channels)
	go func() {
		err := scn.Scan()
		if err != nil {
			panic(err)
		}
	}()

	return nil
}

func (eng *Engine) Record() error {
	err, chId := eng.channels.NewChannel()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	err, stopCb := eng.channels[chId].Play(512)
	if err != nil {
		panic(err)
	}

	go func() {
		defer wg.Done()
		defer func() {
			fmt.Println("STOPPING FIRST")
			stopCb()
		}()
		midi.MidiSim(eng.scanChan, int(chId))
	}()

	return nil
}

type Params struct {
	master     int
	sampleRate int
	chunkSize  int
}

func NewDefaultParams() *Params {
	return &Params{
		master:     0,
		sampleRate: constants.SampleRate,
		chunkSize:  512,
	}
}
