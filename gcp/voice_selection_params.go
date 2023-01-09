package gcp

type VoiceSelectionParams struct {
	LanguageCode string      `json:"languageCode"`
	Name         string      `json:"name"`
	Gender       VoiceGender `json:"ssmlGender"`
}

func NewVoice(params ...string) *VoiceSelectionParams {
	name, locale, gender := "", "", ""
	if len(params) > 0 {
		name = params[0]
	}
	if len(params) > 1 {
		locale = params[1]
	}
	if len(params) > 2 {
		gender = params[2]
	}
	return &VoiceSelectionParams{
		LanguageCode: locale,
		Name:         name,
		Gender:       ParseVoiceGender(gender),
	}
}
