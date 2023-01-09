package gcp

type SynthesizeRequest struct {
	Input       *SynthesisInput       `json:"input"`
	Voice       *VoiceSelectionParams `json:"voice"`
	AudioConfig *AudioConfig          `json:"audioConfig"`
}

func NewTextSynthesizeRequest(text string, voice *VoiceSelectionParams) *SynthesizeRequest {
	return &SynthesizeRequest{
		Input:       NewTextSynthesisInput(text),
		Voice:       voice,
		AudioConfig: NewDefaultAudioConfig(),
	}
}

func NewSSMLSynthesizeRequest(ssml string, voice *VoiceSelectionParams) *SynthesizeRequest {
	return &SynthesizeRequest{
		Input:       NewSSMLSynthesisInput(ssml),
		Voice:       voice,
		AudioConfig: NewDefaultAudioConfig(),
	}
}
