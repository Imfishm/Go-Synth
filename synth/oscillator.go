// This file should implement a basic oscillator.
// It should generate simple waveforms such as sine, square, saw, and triangle,
// with controls for frequency, waveform selection, phase, and detune.

// TODO: An oscillator can likely be better represented as a struct, so that should be worked on.
package synth

import (
	"fmt"
	"math"
)

type Oscillator struct {
	PhaseBegin     float64
	PhaseEnd       float64
	SampleRate     float64
	Frequency      float64
	PhaseIncrement float64
	Phase          float64
}

func (o *Oscillator) GenerateWaveform(wave string) {
	var amplitude float64
	for i := range int(BUFFER_SIZE) {
		switch wave {
		case "sine":
			amplitude = o.GenerateSine()
		case "sawtooth":
			amplitude = o.GenerateSawtooth()
		case "square":
			amplitude = o.GenerateSquare()
		case "triangle":
			amplitude = o.GenerateTriangle()
		default:
			fmt.Printf("Failed to generate %s waveform", wave)
			return
		}
		SampleBuffer[i] = amplitude
		o.AccumulatePhase()
	}
}

func (o *Oscillator) AccumulatePhase() {
	o.Phase += o.PhaseIncrement
	if o.Phase >= o.PhaseEnd {
		o.Phase = o.PhaseBegin
	}
	fmt.Printf("Current Phase: %f\n", o.Phase)
}

// Each of these functions generates a sample based
// on the type of waveform producing the sample

// TODO: Make these functions independent of the oscillator;
// oscillator should call these from a different file, but same package
func (o *Oscillator) GenerateSine() float64 {
	angle := o.Phase / (2 * math.Pi)
	amplitude := math.Sin(angle)
	return amplitude
}

func (o *Oscillator) GenerateSawtooth() float64 {
	amplitude := (o.Phase * 2) - 1
	return amplitude
}

func (o *Oscillator) GenerateSquare() float64 {
	var amplitude float64
	if o.Phase < 0.5 {
		amplitude = -1.0
	} else {
		amplitude = 1.0
	}
	return amplitude
}

func (o *Oscillator) GenerateTriangle() float64 {
	var amplitude float64
	if o.Phase < 0.5 {
		amplitude = (4.0 * o.Phase) - 1
	} else {
		amplitude = (-4.0*o.Phase + 3)
	}
	return amplitude
}
