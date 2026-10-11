// This file should handle audio output for the synthesizer.
// It should provide the audio callback/buffer generation that continuously
// asks the synth engine for samples and sends those samples to the sound device.
package synth

const BUFFER_SIZE uint = 256 // TODO: Needs to be replaced by audio buff size

var SampleBuffer [BUFFER_SIZE]float64

type Audio struct {
	// TODO: Add references needed from the Synth engine
	SampleRate float64
	Channels   int
	BufferSize uint
	Player     any // TODO: decide on audio library; look into Oto
}

// constructor
func NewAudio() *Audio {
	return nil
}

func (a *Audio) Start() {
}

func (a *Audio) Stop() {
}

func (a *Audio) Callback() {
}
