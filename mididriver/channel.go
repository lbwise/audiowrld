package mididriver

import (
	"errors"
	"fmt"
	"time"

	"github.com/lbwise/audiowrld/audio"
)

// So the idea is we receive all these events from the scanner as instructions
// but the channel is always generating samples, but changes based on the instructions
// it receives from the midi scanner

const MaxChannels = 16

type Channel struct {
	Id           uint8
	curPianoRoll []int
	chunkSize    int
	// You would add global midi control properties here
	// instrument oscillator.Instrument
	// track        *engine.InputTrack
}

func (ch *Channel) Play(chunkSize int) (error, func() error) {
	recording := true
	stopCb := func() error {
		fmt.Println("STOPPING")
		recording = false
		return nil
	}

	go func() {
		for recording {
			// Generate chunk
			time.Sleep(time.Second * time.Duration(float64(ch.chunkSize)/float64(audio.SampleRate))) // ~11ms
		}
	}()
	return nil, stopCb
}

func (ch *Channel) Tick(buf audio.Buffer) (error, audio.Buffer) {
	// Generate sound from piano roll
	// generate tick size worth of sound
	return nil, []float32{}
}

func (ch *Channel) NewMsg(msg Msg) {
	if msg.On {
		ch.curPianoRoll[msg.Note] = msg.Velocity
	} else {
		ch.curPianoRoll[msg.Note] = 0
	}
}

type Channels []*Channel

func (m Channels) NewChannel(chunkSize int) (error, uint8) {
	for i, ch := range m {
		if ch == nil {
			m[i] = &Channel{
				Id:           uint8(i),
				curPianoRoll: make([]int, 128),
				chunkSize:    chunkSize,
				//  instrument: inst,
			}
			return nil, uint8(i)
		}
	}
	return errors.New("too many channels"), 0
}
