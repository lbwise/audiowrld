package io

import (
	"os"

	"github.com/gordonklaus/portaudio"
)

func InitDriver() func() {

	f, err := os.Open("../exports/test-audio.wav")
	if err != nil {
		panic(err)
	}

	id, err := portaudio.Initialize()

	out := make([]int32, 8192)
	stream, err := portaudio.OpenDefaultStream(0, 1, 44100, len(out), &out)
	if err != nil {
		panic(err)
	}

	err = stream.Start()
	if err != nil {
		panic(err)
	}

	return func() {
		portaudio.Terminate()
		stream.Close()
		stream.Stop()
	}

}
