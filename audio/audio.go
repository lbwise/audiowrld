package audio

const MaxAmp int = 32760
const SampleRate int = 44100

func NewBuffer() Buffer { return Buffer{} }

type Buffer []float32

type Params struct {
	master     int
	sampleRate int
	chunkSize  int
}

func NewDefaultParams() *Params {
	return &Params{
		master:     0,
		sampleRate: SampleRate,
		chunkSize:  512,
	}
}
