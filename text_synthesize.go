package tts_integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	jsonContentType = "Content-Type: application/json"
)

type TextSynthesizeResponse struct {
	AudioContent string `json:"audioContent"`
}

func PostTextSynthesize(hc *http.Client, txt string, voice *VoiceSelectionParams, key string) (*TextSynthesizeResponse, error) {

	tsu := TextSynthesizeUrl(key)
	sr := NewSynthesizeRequest(txt, voice)

	req, err := json.Marshal(sr)
	if err != nil {
		return nil, err
	}

	fmt.Printf(string(req))

	resp, err := hc.Post(tsu.String(), jsonContentType, bytes.NewReader(req))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tsr *TextSynthesizeResponse
	err = json.NewDecoder(resp.Body).Decode(&tsr)

	return tsr, err
}
