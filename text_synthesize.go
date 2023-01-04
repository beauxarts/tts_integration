package google_tts_integration

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
)

const (
	jsonContentType = "Content-Type: application/json"
)

type TextSynthesizeResponse struct {
	AudioContent string `json:"audioContent"`
}

func (tsr *TextSynthesizeResponse) Bytes() ([]byte, error) {
	return base64.StdEncoding.DecodeString(tsr.AudioContent)
}

func PostTextSynthesize(hc *http.Client, txt string, voice *VoiceSelectionParams, key string) (*TextSynthesizeResponse, error) {

	tsu := TextSynthesizeUrl(key)
	sr := NewSynthesizeRequest(txt, voice)

	req, err := json.Marshal(sr)
	if err != nil {
		return nil, err
	}

	resp, err := hc.Post(tsu.String(), jsonContentType, bytes.NewReader(req))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New("text:synthesize error status " + resp.Status)
	}

	var tsr *TextSynthesizeResponse
	err = json.NewDecoder(resp.Body).Decode(&tsr)

	return tsr, err
}
