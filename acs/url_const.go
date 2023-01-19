package acs

// /cognitiveservices/voices/list
const (
	httpsScheme                            = "https"
	ttsSpeechMicrosoftHost                 = "tts.speech.microsoft.com"
	regionalTtsSpeechMicrosoftHostTemplate = "{region}." + ttsSpeechMicrosoftHost
	cognitiveServicesPath                  = "/cognitiveservices"
	voicesListPath                         = cognitiveServicesPath + "/voices/list"
	textToSpeechV1Path                     = cognitiveServicesPath + "/v1"
)

const (
	OcpApimSubscriptionKeyHeader = "Ocp-Apim-Subscription-Key"
)
