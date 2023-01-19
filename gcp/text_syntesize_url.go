package gcp

import "net/url"

func TextSynthesizeUrl() *url.URL {
	return &url.URL{
		Scheme: httpsScheme,
		Host:   textToSpeechAPIHost,
		Path:   textSynthesizePath,
	}
}
