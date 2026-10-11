// This file should implement the synthesizer's simple wave shaper.
// It should take an audio sample and transform its amplitude to add a basic
// distortion/saturation-style character without implementing full effects.
package synth

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

func (w *WaveshaperType) Process() {
}

func (w *WaveshaperType) SetDrive() {
}

func (w *WaveshaperType) SetCurve() {
}

func (w *WaveshaperType) SetGain() {
}
