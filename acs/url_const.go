package acs

// /cognitiveservices/voices/list
const (
	httpsScheme                               = "https"
	ttsSpeechMicrosoftHost                    = "tts.speech.microsoft.com"
	regionalTtsSpeechMicrosoftHostTemplate    = "{region}." + ttsSpeechMicrosoftHost
	apiCognitiveMicrosoftHost                 = "api.cognitive.microsoft.com"
	regionalApiCognitiveMicrosoftHostTemplate = "{region}." + apiCognitiveMicrosoftHost
	cognitiveServicesPath                     = "/cognitiveservices"
	voicesListPath                            = cognitiveServicesPath + "/voices/list"
	textToSpeechV1Path                        = cognitiveServicesPath + "/v1"
	issueTokenPath                            = "/sts/v1.0/issueToken"
)

const (
	OcpApimSubscriptionKeyHeader = "Ocp-Apim-Subscription-Key"
)
