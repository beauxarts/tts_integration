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

type Synthesizer struct {
	httpClient *http.Client
	voice      *VoiceSelectionParams
	key        string
}

func NewSynthesizer(hc *http.Client, voice *VoiceSelectionParams, key string) *Synthesizer {
	return &Synthesizer{
		httpClient: hc,
		voice:      voice,
		key:        key,
	}
}

func (s *Synthesizer) PostText(text string) (*TextSynthesizeResponse, error) {
	return s.postSynthesizeRequest(text, Text)
}

func (s *Synthesizer) PostSSML(ssml string) (*TextSynthesizeResponse, error) {
	return s.postSynthesizeRequest(ssml, SSML)
}

func (s *Synthesizer) postSynthesizeRequest(content string, contentType SynthesisInputType) (*TextSynthesizeResponse, error) {

	var newRequest func(string, *VoiceSelectionParams) *SynthesizeRequest
	switch contentType {
	case Text:
		newRequest = NewTextSynthesizeRequest
	case SSML:
		newRequest = NewSSMLSynthesizeRequest
	}

	var sr *SynthesizeRequest
	if newRequest != nil {
		sr = newRequest(content, s.voice)
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

	tsu := TextSynthesizeUrl(s.key)

	resp, err := s.httpClient.Post(tsu.String(), jsonContentType, bytes.NewReader(req))
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
