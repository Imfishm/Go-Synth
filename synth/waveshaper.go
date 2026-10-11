package synth

import (
	"math"
)

type WaveshaperType struct {
	Drive float64
	Gain  float64
	Curve CurveType
}

type CurveType int

const (
	Softclip CurveType = iota
	Hardclip
	Wavefold
)

func (w *WaveshaperType) Process(sample float64) float64 {
	sample *= w.Drive

	switch w.Curve {
	case Softclip:
		sample = math.Tanh(sample)

	case Hardclip:
		if sample > 1.0 {
			sample = 1.0
		}
		if sample < -1.0 {
			sample = -1.0
		}

	case Wavefold:
		for sample > 1.0 || sample < -1.0 {
			if sample > 1.0 {
				sample = 2.0 - sample
			}
			if sample < -1.0 {
				sample = -2.0 - sample
			}
		}

	default:
		// Leave sample unchanged
	}

	return sample * w.Gain
}

func (w *WaveshaperType) SetDrive(d float64) {
	w.Drive = d
}

func (w *WaveshaperType) SetGain(g float64) {
	w.Gain = g
}

func (w *WaveshaperType) SetCurve(c CurveType) {
	w.Curve = c
}
