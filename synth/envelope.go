// This file should implement a simple ADSR envelope.
// It should handle attack, decay, sustain, and release stages and provide an
// amplitude value that changes over time for a note.
package synth

type Envelope struct {
	Attack                float64
	Decay                 float64
	Sustain               float64
	Release               float64
	Amplitude             float64
	ReleaseStartAmplitude float64
	SampleRate            float64
	Progress              int // in # of samples
	Stage                 State
}

// State enum
type State int

const (
	Idle State = iota
	Attack
	Decay
	Sustain
	Release
)

func (e Envelope) Trigger() {
}

func (e Envelope) NoteOff() {
}

func (e Envelope) Process() {
}

func (e Envelope) IsActive() {
}
