// This file should contain the main synthesizer engine.
// It should combine oscillators, envelopes, LFOs, detuning, and the wave shaper
// into the final audio signal produced by the synthesizer.
package synth

const MAX_POLYPHONY int = 100
const EVENT_BUFFER int = 300

type Synth struct {
	VoicePool     [MAX_POLYPHONY]Voice
	SampleRate    float64
	PatchSettings any // needs to be a custom type
	Detune        int // in cents
	Waveshaper    WaveshaperType
	MasterVolume  float64
	EventQueue    chan Note // initialized buffer of size EVENT_BUFFER
}

// constructor
func NewSynth() *Synth {
	return nil
}

func (s *Synth) Process() {
}

func (s *Synth) NoteOn() {
}

func (s *Synth) NoteOff() {
}

func (s *Synth) FillBuffer() {
}

func (s *Synth) Panic() { // triggers when all notes are off
}

func (s *Synth) SetPatch() {
}

func (s *Synth) SetWaveform() {
}

func (s *Synth) SetAttack() {
}

func (s *Synth) SetDecay() {
}

func (s *Synth) SetRelease() {
}

func (s *Synth) SetSustain() {
}

func (s *Synth) SetLFORate() {
}

func (s *Synth) SetLFOShape() {
}

func (s *Synth) SetLFODepth() {
}

func (s *Synth) SetLFOTarget() {
}

func (s *Synth) SetDetune() {
}

func (s *Synth) SetDrive() {
}

func (s *Synth) SetMaster() {
}
