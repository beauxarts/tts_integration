package gcp

import "net/url"

const (
	languageCodeParam = "languageCode"
)

func VoicesUrl(langCode string) *url.URL {

	if langCode == "" {
		return nil
	}

	u := &url.URL{
		Scheme: httpsScheme,
		Host:   textToSpeechAPIHost,
		Path:   voicesPath,
	}

	q := u.Query()
	q.Set(languageCodeParam, langCode)
	u.RawQuery = q.Encode()

	return u
}
