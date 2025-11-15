package engine

import (
	"github.com/lbwise/audiowrld/audio"
)

func Mix(buf audio.Buffer, chs []audio.Buffer) (error, audio.Buffer) {
	for i := 0; i < len(buf); i++ {
		for chId, chBuf := range chs {
			// Normalise and mix
		}
		buf[i] = 1
	}
	return nil, buf
}
