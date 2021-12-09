package dsp

func L16MixSat160(dst, src *int16) {
	panic("Not implemented")
}

// Compresses a PCM audio sample into a G.711 μ-Law sample. The BSR instruction
// is what makes this code fast.
//
// TODO(jart): How do I make assembly use proper types?
func LinearToUlaw(linear int64) (ulaw int64) {
	panic("Not implemented")

}

// Turns a μ-Law byte back into an audio sample.
func UlawToLinear(ulaw int64) (linear int64) {
	panic("Not implemented")

}