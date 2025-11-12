package mididriver

import (
	"errors"
)

type Msg struct {
	On       bool
	Channel  int
	Note     int
	Velocity int
}

// NewMsg converts raw midi msg to Msg obj
func NewMsg(rawMsg []byte) (error, *Msg) {
	if len(rawMsg) != 3 {
		return errors.New("invalid raw message"), nil
	}

	note := rawMsg[0]
	velocity := rawMsg[1]
	status := rawMsg[2]

	if !validMidiParam(note) {
		return errors.New("invalid note value"), nil
	} else if !validMidiParam(velocity) {
		return errors.New("invalid velocity value"), nil
	}

	err, on := midiOn(status)
	if err != nil {
		return err, nil
	}

	err, ch := midiChannel(status)
	return nil, &Msg{Note: int(note), Velocity: int(velocity), Channel: ch, On: on}
}

// NewRawMsg is used to synthesize midi messages
func NewRawMsg(note, velocity, ch int, on bool) []byte {
	var status byte
	if on {
		status = 0x90 | byte(ch)
	} else {
		status = 0x80 | byte(ch)
	}
	return []byte{byte(note), byte(velocity), status}
}

func midiOn(status byte) (error, bool) {
	var on bool
	if status&0xF0 == 0x90 {
		on = true
	} else if status&0xF0 == 0x80 {
		on = false
	} else {
		return errors.New("invalid status"), false
	}
	return nil, on
}

func midiChannel(status byte) (error, int) {
	ch := int(status & 0x0F)
	if ch < 0 || ch >= MaxChannels {
		return errors.New("invalid channel status"), 0
	}
	return nil, ch
}
