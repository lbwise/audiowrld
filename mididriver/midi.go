package mididriver

import (
	"fmt"
	"time"
)

type RawMsg []byte

func NewScanner(in <-chan RawMsg, channels Channels) *Scanner {
	return &Scanner{In: in, channels: channels}
}

type Scanner struct {
	In       <-chan RawMsg
	channels Channels
}

func (scanner *Scanner) Scan() error {
	for msg := range scanner.In {
		fmt.Println(time.Now(), fmt.Sprintf("note: %b, velocity: %b, status: %b", msg[0], msg[1], msg[2]))
		err, msg := NewMsg(msg) // should not return pointer
		if err != nil {
			return err
		}

		for _, ch := range scanner.channels {
			if ch != nil {
				ch.NewEvent(*msg)
			}
		}
	}
	return nil
}

func validMidiParam(value byte) bool {
	return value <= 127
}

func MidiSim(out chan<- RawMsg, chId int) {
	data := NewRawMsg(60, 100, chId, true)
	for i := 0; i < 10; i++ {
		out <- data
		fmt.Println("MSG SENT")
		time.Sleep(time.Millisecond * 1000) // simulate timing
	}
	close(out)
}
