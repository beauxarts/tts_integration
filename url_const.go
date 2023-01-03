package tts_integration

const (
	httpsScheme         = "https"
	googleAPIsHost      = "googleapis.com"
	textToSpeechAPIHost = "texttospeech." + googleAPIsHost
	v1Path              = "/v1"
	voicesPath          = v1Path + "/voices"
	textSynthesizePath  = v1Path + "/text:synthesize"
)
