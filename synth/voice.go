// This file should represent one currently playing synthesizer voice.
// It should keep track of a note's frequency, oscillator state, envelopes,
// LFO state, and whether the voice is active or releasing.
package synth

type Voice struct {
	Frequency       float64
	OscillatorState any
	Envelopes       any // Starting with just volume envelope
	LfoState        any
	Releasing       bool
}

// All functions are prototypes currently; not usable

func (v Voice) Start(n Note) {
}

func (v Voice) Release() {
}

func (v Voice) GenerateAudio() {
}

func (v Voice) IsAvailable() bool {
	return false
}

func (v Voice) IsActive() bool {
	return false
}
