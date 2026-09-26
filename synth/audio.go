// This file should handle audio output for the synthesizer.
// It should provide the audio callback/buffer generation that continuously
// asks the synth engine for samples and sends those samples to the sound device.
package synth

const BufferSize uint = 256

var SampleBuffer [BufferSize]float64
