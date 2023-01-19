package gcp

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/beauxarts/tts_integration"
	"net/http"
)

func TextSynthesize(hc *http.Client, content string, voice *VoiceSelectionParams, contentType tts_integration.SynthesisInputType, key string) (*TextSynthesizeResponse, error) {

	var newRequest func(string, *VoiceSelectionParams) *SynthesizeRequest
	switch contentType {
	case tts_integration.Text:
		newRequest = NewTextSynthesizeRequest
	case tts_integration.SSML:
		newRequest = NewSSMLSynthesizeRequest
	}

	var sr *SynthesizeRequest
	if newRequest != nil {
		sr = newRequest(content, voice)
	} else {
		return nil, errors.New("unknown content type " + contentType.String())
	}

	if sr == nil {
		return nil, errors.New("error creating synthesize request")
	}

	req, err := json.Marshal(sr)
	if err != nil {
		return nil, err
	}

	tsu := TextSynthesizeUrl()

	tsReq, err := http.NewRequest(http.MethodPost, tsu.String(), bytes.NewReader(req))
	if err != nil {
		return nil, err
	}

	tsReq.Header.Add("X-goog-api-key", key)

	resp, err := hc.Do(tsReq)
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
