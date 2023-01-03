package tts_integration

import "net/url"

const (
	languageCodeParam = "languageCode"
	keyParam          = "key"
)

func VoicesUrl(lc, key string) *url.URL {

	if lc == "" || key == "" {
		return nil
	}

	u := &url.URL{
		Scheme: httpsScheme,
		Host:   textToSpeechAPIHost,
		Path:   voicesPath,
	}

	q := u.Query()
	q.Set(languageCodeParam, lc)
	q.Set(keyParam, key)
	u.RawQuery = q.Encode()

	return u
}
