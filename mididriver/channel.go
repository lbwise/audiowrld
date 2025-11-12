package mididriver

import (
	"errors"
	"fmt"
	"time"

	"github.com/lbwise/audiowrld/simplesynth/constants"
)

// So the idea is we receive all these events from the scanner as instructions
// but the channel is always generating samples, but changes based on the instructions
// it receives from the midi scanner

const MaxChannels = 16

type Channel struct {
	Id uint8
	//Instrument oscillator.Instrument
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
			time.Sleep(time.Second * time.Duration(float64(chunkSize)/float64(constants.SampleRate)))
		}
	}()
	return nil, stopCb
}

func (ch *Channel) NewEvent(msg Msg) {
	fmt.Println("MSG:", msg)
}

type Channels []*Channel

func (m Channels) NewChannel() (error, uint8) {
	for i, ch := range m {
		if ch == nil {
			//ch := &Channel{Id: id, Instrument: inst}
			m[i] = &Channel{Id: uint8(i)}
			return nil, uint8(i)
		}
	}
	return errors.New("too many channels"), 0
}
