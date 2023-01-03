package tts_integration

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
