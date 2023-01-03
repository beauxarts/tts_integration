package tts_integration

type VoiceSelectionParams struct {
	LanguageCode string      `json:"languageCode"`
	Name         string      `json:"name"`
	Gender       VoiceGender `json:"ssmlGender"`
}

func NewVoiceSelectionParams(lc, nm string, gr VoiceGender) *VoiceSelectionParams {
	return &VoiceSelectionParams{
		LanguageCode: lc,
		Name:         nm,
		Gender:       gr,
	}
}
