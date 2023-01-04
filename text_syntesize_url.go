package google_tts_integration

import "net/url"

func TextSynthesizeUrl(key string) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   textToSpeechAPIHost,
		Path:   textSynthesizePath,
	}

	q := u.Query()
	q.Set(keyParam, key)
	u.RawQuery = q.Encode()

	return u
}
