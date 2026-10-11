// This file should define the basic representation of a musical note.
// It should provide the information needed to start and stop a note, including
// MIDI-style note number and velocity information.
package synth

import "math"

type Note struct {
	Number   uint8
	Velocity uint8
	Channel  uint8
	Gate     bool
	// TODO: Add timestamp for sample accurate timing later on (not needed rn)
}

func (n *Note) NoteToFrequency(noteNum uint8) float64 {
	return 440.0 * math.Pow(2, ((float64(noteNum)-69)/12))
}
