package tts_integration

import (
	"io"
	"time"
)

type Synthesizer interface {
	IsSSMLSupported() bool
	IsWriterRequired() bool
	IsNameRequired() bool
	VoicesStrings(params ...string) ([]string, error)
	WriteText(t string, w io.Writer, n string) error
	WriteSSML(s string, w io.Writer, n string) error
	DecorateWithPauses(text string, d time.Duration) (string, SynthesisInputType)
}
