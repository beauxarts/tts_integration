package acs

import (
	"net/url"
	"strings"
)

func TextToSpeechUrl(region string) *url.URL {
	return &url.URL{
		Scheme: httpsScheme,
		Host:   strings.Replace(regionalTtsSpeechMicrosoftHostTemplate, "{region}", region, -1),
		Path:   textToSpeechV1Path,
	}
}
