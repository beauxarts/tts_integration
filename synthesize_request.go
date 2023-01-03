package tts_integration

type SynthesizeRequest struct {
	Input       *SynthesisInput       `json:"input"`
	Voice       *VoiceSelectionParams `json:"voice"`
	AudioConfig *AudioConfig          `json:"audioConfig"`
}

func NewSynthesizeRequest(text string, voice *VoiceSelectionParams) *SynthesizeRequest {
	return &SynthesizeRequest{
		Input:       NewTextSynthesisInput(text),
		Voice:       voice,
		AudioConfig: NewDefaultAudioConfig(),
	}
}
