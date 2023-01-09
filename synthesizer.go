package tts_integration

import "io"

type Synthesizer interface {
	SynthesizeText(t string, w io.Writer) error
	SynthesizeSSML(s string, w io.Writer) error
	VoicesStrings(params ...string) ([]string, error)
}
