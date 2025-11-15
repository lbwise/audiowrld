package engine

import (
	"sync"

	"github.com/lbwise/audiowrld/io"
	midi "github.com/lbwise/audiowrld/mididriver"
	"github.com/lbwise/audiowrld/simplesynth/constants"
	inst "github.com/lbwise/audiowrld/simplesynth/oscillator"
)

type Ticker interface {
	Tick(in AudioBuffer) (error, AudioBuffer)
}

type Clock int

type Engine struct {
	params       *Params
	channels     midi.Channels
	instruments  []inst.Instrument
	outputDevice *io.OutputDevice
	scanChan     chan midi.RawMsg
	tracks       []Track
	master       Track
	tick         int
	outputBuffer []AudioBuffer
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
			stopCb()
		}()
		midi.MidiSim(eng.scanChan, int(chId))
	}()

	return nil
}

func (eng *Engine) Start(checkStop func() bool) error {

	eng.scanChan = make(chan midi.RawMsg, 64)
	scn := midi.NewScanner(eng.scanChan, eng.channels)
	go func() {
		err := scn.Scan()
		if err != nil {
			panic(err)
		}
	}()

}

func (eng *Engine) Tick() {
	// The tick order should be something like this
	eng.tick++

	for _, track := range eng.tracks {
		buf := make([]AudioBuffer, eng.params.chunkSize)
		if track.Type() == MidiTrack {
			midiTr := InputTrack(track)
			err, buf := midiTr.channel.Tick(buf)
			midiTr.Tick(buf)

		}

	}

	for _, ch := range eng.channels {
		if ch != nil {
			err, buf = ch.Tick()
			if err != nil {
				panic(err)
			}
		}
	}

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
