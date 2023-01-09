package tts_integration

type SynthesisInputType string

const (
	Text SynthesisInputType = "Text"
	SSML SynthesisInputType = "SSML"
)

func (sit SynthesisInputType) String() string {
	return string(sit)
}
