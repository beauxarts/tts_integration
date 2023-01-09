package gcp

type SynthesisInputType string

const (
	Text SynthesisInputType = "Text"
	SSML SynthesisInputType = "SSML"
)

func (sit SynthesisInputType) String() string {
	return string(sit)
}

type SynthesisInput struct {
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
}

func NewTextSynthesisInput(text string) *SynthesisInput {
	return &SynthesisInput{
		Text: text,
	}
}

func NewSSMLSynthesisInput(ssml string) *SynthesisInput {
	return &SynthesisInput{
		SSML: ssml,
	}
}
