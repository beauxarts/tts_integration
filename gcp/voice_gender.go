package gcp

import "strings"

type VoiceGender string

const (
	//VoiceGenderUnspecified - An unspecified gender. In VoiceSelectionParams, this means that the client doesn't care which gender the selected voice will have. In the Voice field of ListVoicesResponse, this may mean that the voice doesn't fit any of the other categories in this enum, or that the gender of the voice isn't known.
	VoiceGenderUnspecified VoiceGender = "SSML_VOICE_GENDER_UNSPECIFIED"
	//Male - A male voice.
	Male = "MALE"
	//Female - A female voice.
	Female = "FEMALE"
	//Neutral - A gender-neutral voice. This voice is not yet supported.
	Neutral = "NEUTRAL"
)

func ParseVoiceGender(gr string) VoiceGender {
	switch strings.ToLower(gr) {
	case strings.ToLower(Male):
		return Male
	case strings.ToLower(Female):
		return Female
	case strings.ToLower(Neutral):
		return Neutral
	}
	return VoiceGenderUnspecified
}
