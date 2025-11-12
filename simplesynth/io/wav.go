package io

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/lbwise/audiowrld/simplesynth/constants"
)

func CreateWAV(dir string, name string, audio []int16, sampleSize int) {
	f, err := os.Create(fmt.Sprintf("%s/%s.wav", dir, name))
	if err != nil {
		panic(fmt.Sprintf("Could not create WAV file: %s", err))
	}

	writeWavHeader(f, sampleSize, constants.SampleRate, 16, 1)

	for _, sample := range audio {
		err = binary.Write(f, binary.LittleEndian, sample)
	}

	if err != nil {
		panic(fmt.Sprintf("Could not write to WAV file: %s", err))
	}
}

func writeWavHeader(f *os.File, numSamples, sampleRate, bitsPerSample, numChannels int) {
	byteRate := sampleRate * numChannels * bitsPerSample / 8
	blockAlign := numChannels * bitsPerSample / 8
	dataSize := numSamples * numChannels * bitsPerSample / 8
	chunkSize := 36 + dataSize

	f.Write([]byte("RIFF"))
	binary.Write(f, binary.LittleEndian, uint32(chunkSize))
	f.Write([]byte("WAVE"))

	// fmt subchunk
	f.Write([]byte("fmt "))
	binary.Write(f, binary.LittleEndian, uint32(16)) // Subchunk1Size
	binary.Write(f, binary.LittleEndian, uint16(1))  // PCM format
	binary.Write(f, binary.LittleEndian, uint16(numChannels))
	binary.Write(f, binary.LittleEndian, uint32(sampleRate))
	binary.Write(f, binary.LittleEndian, uint32(byteRate))
	binary.Write(f, binary.LittleEndian, uint16(blockAlign))
	binary.Write(f, binary.LittleEndian, uint16(bitsPerSample))

	// data subchunk
	f.Write([]byte("data"))
	binary.Write(f, binary.LittleEndian, uint32(dataSize))
}
