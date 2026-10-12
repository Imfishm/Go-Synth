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

func (l *LFO) Process() float64 {
	out := l.Value()

	phase := l.Phase + (l.Rate / l.SampleRate)

	if phase >= 1 {
		phase = phase - 1.0
		// TODO: Add conditional for if waveshape is sample and hold
	}

	return out
}

func (l *LFO) Reset() {
	l.Phase = 0.0
}

func (l *LFO) Value() float64 {
	// TODO: Move wave funcs from osc into their own file
	// and use those generators for lfo gen here

	switch l.Waveshape {
	default:
		return 0.0
	}
}

func (l *LFO) SetRate(r float64) {
	l.Rate = r
}

func (l *LFO) SetWaveshape(w WaveformShape) {
	l.Waveshape = w
}
