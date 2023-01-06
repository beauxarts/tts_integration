package google_tts_integration

type AudioConfig struct {
	//Encoding - Required. The format of the audio byte stream.
	Encoding AudioEncoding `json:"audioEncoding"`
	//SpeakingRate - Optional. Input only. Speaking rate/speed, in the range [0.25, 4.0]. 1.0 is the normal native speed supported by the specific voice. 2.0 is twice as fast, and 0.5 is half as fast. If unset(0.0), defaults to the native 1.0 speed. Any other values < 0.25 or > 4.0 will return an error.
	SpeakingRate float32 `json:"speakingRate,omitempty"`
	//Pitch - Optional. Input only. Speaking pitch, in the range [-20.0, 20.0]. 20 means increase 20 semitones from the original pitch. -20 means decrease 20 semitones from the original pitch.
	Pitch float32 `json:"pitch,omitempty"`
	//VolumeGainDb - Optional. Input only. Volume gain (in dB) of the normal native volume supported by the specific voice, in the range [-96.0, 16.0]. If unset, or set to a value of 0.0 (dB), will play at normal native signal amplitude. A value of -6.0 (dB) will play at approximately half the amplitude of the normal native signal amplitude. A value of +6.0 (dB) will play at approximately twice the amplitude of the normal native signal amplitude. Strongly recommend not to exceed +10 (dB) as there's usually no effective increase in loudness for any value greater than that.
	VolumeGainDb float32 `json:"volumeGainDb,omitempty"`
	//SampleRateHertz - Optional. The synthesis sample rate (in hertz) for this audio. When this is specified in SynthesizeSpeechRequest, if this is different from the voice's natural sample rate, then the synthesizer will honor this request by converting to the desired sample rate (which might result in worse audio quality), unless the specified sample rate is not supported for the encoding chosen, in which case it will fail the request and return google.rpc.Code.INVALID_ARGUMENT.
	SampleRateHertz  int            `json:"sampleRateHertz,omitempty"`
	EffectsProfileId []AudioProfile `json:"effectsProfileId,omitempty"`
}

func NewDefaultAudioConfig() *AudioConfig {
	return &AudioConfig{
		Encoding:         OggOpus,
		SpeakingRate:     1.0,
		Pitch:            0.0,
		VolumeGainDb:     0.0,
		SampleRateHertz:  0,
		EffectsProfileId: []AudioProfile{HeadphoneClassDevice},
	}
}

const DefaultEncodingExt = ".ogg"
