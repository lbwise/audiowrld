package mididriver

import (
	"errors"

	"github.com/lbwise/audiowrld/simplesynth/oscillator"
)

const MAX_CHANNELS = 16

type Channels []*Channel

func (m Channels) NewMidiChannel(inst oscillator.Instrument) (error, *Channel) {
	if len(m) == MAX_CHANNELS-1 {
		return errors.New("too many channels"), nil
	}
	id := len(m) // Is this the best way to do the id?
	ch := &Channel{Id: id, Instrument: inst}
	m[id] = ch
	return nil, ch

}

type Channel struct {
	Id         int
	Instrument oscillator.Instrument
}

func NewMidiRawMsg(note, velocity, ch int, channels []MidiChannel) (error, *MidiRawMsg) {
	if !isOneByte(note) {
		return errors.New("MidiRawMsg note invalid"), nil
	} else if !isOneByte(velocity) {
		return errors.New("MidiRawMsg velocity invalid"), nil
	} else if !isValidChNum(ch, channels) {
		return errors.New("MidiRawMsg channel number invalid"), nil
	}
	return nil, &MidiRawMsg{Note: uint8(note), Velocity: uint8(velocity), Channel: uint8(ch)}
}

type MidiRawMsg struct {
	Note     uint8
	Velocity uint8
	Channel  uint8
}

func isValidChNum(chNum int, channels []MidiChannel) bool {
	return chNum >= 0 && chNum <= len(channels)
}

func isOneByte(value int) bool {
	return value >= 0 && value <= 127
}
