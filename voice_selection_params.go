package google_tts_integration

type VoiceSelectionParams struct {
	LanguageCode string      `json:"languageCode"`
	Name         string      `json:"name"`
	Gender       VoiceGender `json:"ssmlGender"`
}

func NewVoice(lc, nm, gr string) *VoiceSelectionParams {
	return &VoiceSelectionParams{
		LanguageCode: lc,
		Name:         nm,
		Gender:       ParseVoiceGender(gr),
	}
}
