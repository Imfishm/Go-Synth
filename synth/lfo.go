// This file should implement a basic low-frequency oscillator.
// It should generate a slow modulation signal that can be used for simple
// pitch, volume, or other parameter modulation.
package synth

type LFO struct {
	Rate       float64
	Phase      float64
	SampleRate float64
	Waveshape  WaveformShape
}

type WaveformShape int

const (
	Sine WaveformShape = iota
	Triangle
	Square
	Sawtooth
	// TODO: Add SampleAndHold field and related methods/fields
)

func (l *LFO) Process() {
}

func (l *LFO) Reset() {
}

func (l *LFO) Value() {
}

func (l *LFO) SetRate() {
}

func (l *LFO) SetShape() {
}
