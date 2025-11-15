package engine

import (
	"errors"
	"fmt"
	//"sync"

	"github.com/lbwise/audiowrld/audio"
	inst "github.com/lbwise/audiowrld/instrument"
	"github.com/lbwise/audiowrld/io"
	midi "github.com/lbwise/audiowrld/mididriver"
)

type Ticker interface {
	Tick(in audio.Buffer) (error, audio.Buffer)
}

type Clock int

type Engine struct {
	params       *audio.Params
	channels     midi.Channels
	instruments  []inst.Instrument
	outputDevice *io.OutputDevice
	scanChan     chan midi.RawMsg
	tracks       []Track
	master       Track
	tick         int
	outputBuffer []audio.Buffer
}

func (eng *Engine) Start(checkStop func() bool) error {

	// Start midi scanner (ONLY IF MIDI CHANNEL APPLICABLE?)
	//var wg sync.WaitGroup
	//wg.Add(2)
	//eng.scanChan = make(chan midi.RawMsg, 64)
	//scn := midi.NewScanner(eng.scanChan, eng.channels)
	//go func() {
	//	err := scn.Scan()
	//	if err != nil {
	//		panic(err)
	//	}
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	defer func() {
	//		stopCb()
	//	}()
	//	midi.MidiSim(eng.scanChan, int(chId))
	//}()

	// This need to be done per callback request by the audio engine
	for {
		masterBuf := audio.NewBuffer()
		err, masterBuf := eng.Tick(masterBuf)
		if err != nil {
			panic(err) // fix this
		}
		//// Save and play audio
		//io.Save
	}
}

func (eng *Engine) Tick(buf audio.Buffer) (error, audio.Buffer) {
	// The tick order should be something like this
	eng.tick++

	// These should all be go routines and then collected to mix
	for _, track := range eng.tracks {
		trackBuf := audio.NewBuffer()
		if track.Type() == MidiTrack {
			midiTr := InputTrack(track)
			err, buf := midiTr.channel.Tick(trackBuf)
			trackBuf = buf
			if err != nil {
				return errors.New(fmt.Sprintf("MIDI TICK ERROR: %s", err)), buf
			}

			//midiTr.Tick(buf)
		} else {
			//audioTrack := AudioTrack(track)
			return nil, trackBuf
		}

		// Processing (or is that done in track.Tick ?)
	}

	// Mix
	err, outBuf := Mix(buf, []audio.Buffer{})
	if err != nil {
		return err, buf
	}

	// Process mix (or is that done in mix track.Tick)

	buf = outBuf
	return nil, buf
}
