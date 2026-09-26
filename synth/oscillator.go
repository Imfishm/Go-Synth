// This file should implement a basic oscillator.
// It should generate simple waveforms such as sine, square, saw, and triangle,
// with controls for frequency, waveform selection, phase, and detune.
package synth

import (
	"fmt"
	"math"
)

const PhaseBegin float64 = 0.0
const PhaseEnd float64 = 1.0

const SampleRate uint = 44100 // Samples per second
const Frequency uint = 440    // Hertz
const PhaseIncrement float64 = float64(Frequency / SampleRate)

var Phase float64 = PhaseBegin

func GenerateWaveform(waveType string) {
	var amplitude float64
	for i := range int(BufferSize) {
		switch waveType {
		case "sine":
			amplitude = GenerateSine()
		case "sawtooth":
			amplitude = GenerateSawtooth()
		case "square":
			amplitude = GenerateSquare()
		case "triangle":
			amplitude = GenerateTriangle()
		default:
			fmt.Printf("\033[31mError: Failed you generate waveform of requested type %v\033[0m\n", waveType)
			return
		}
		SampleBuffer[i] = amplitude
		AccumulatePhase()
	}
}

func AccumulatePhase() {
	for {
		Phase += PhaseIncrement
		if Phase >= PhaseEnd {
			Phase = PhaseBegin
		}
		fmt.Printf("Current Phase: %f\n", Phase)
	}
}

// Each of these functions generates a sample based
// on the type of waveform producing the sample

func GenerateSine() float64 {
	angle := Phase / (2 * math.Pi)
	amplitude := math.Sin(angle)
	return amplitude
}

func GenerateSawtooth() float64 {
	amplitude := (Phase * 2) - 1
	return amplitude
}

func GenerateSquare() float64 {
	var amplitude float64
	if Phase < 0.5 {
		amplitude = -1.0
	} else {
		amplitude = 1.0
	}
	return amplitude
}

func GenerateTriangle() float64 {
	var amplitude float64
	if Phase < 0.5 {
		amplitude = (4.0 * Phase) - 1
	} else {
		amplitude = (-4.0*Phase + 3)
	}
	return amplitude
}
