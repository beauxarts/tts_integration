package gcp

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beauxarts/tts_integration"
	"io"
	"net/http"
	"time"
)

const (
	jsonContentType = "Content-Type: application/json"
)

type textSynthesizeResponse struct {
	AudioContent string `json:"audioContent"`
}

func (tsr *textSynthesizeResponse) Bytes() ([]byte, error) {
	return base64.StdEncoding.DecodeString(tsr.AudioContent)
}

type Synthesizer struct {
	httpClient *http.Client
	voice      *VoiceSelectionParams
	key        string
}

func NewSynthesizer(hc *http.Client, key string, voiceParams ...string) tts_integration.Synthesizer {
	return &Synthesizer{
		httpClient: hc,
		voice:      NewVoice(voiceParams...),
		key:        key,
	}
}

func (s *Synthesizer) postText(text string) (*textSynthesizeResponse, error) {
	return s.postSynthesizeRequest(text, tts_integration.Text)
}

func (s *Synthesizer) postSSML(ssml string) (*textSynthesizeResponse, error) {
	return s.postSynthesizeRequest(ssml, tts_integration.SSML)
}

func (s *Synthesizer) postSynthesizeRequest(content string, contentType tts_integration.SynthesisInputType) (*textSynthesizeResponse, error) {

	var newRequest func(string, *VoiceSelectionParams) *SynthesizeRequest
	switch contentType {
	case tts_integration.Text:
		newRequest = NewTextSynthesizeRequest
	case tts_integration.SSML:
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

	var tsr *textSynthesizeResponse
	err = json.NewDecoder(resp.Body).Decode(&tsr)

	return tsr, err
}

func (s *Synthesizer) synthesize(content string, contentType tts_integration.SynthesisInputType, w io.Writer) error {

	var post func(string) (*textSynthesizeResponse, error)
	switch contentType {
	case tts_integration.Text:
		post = s.postText
	case tts_integration.SSML:
		post = s.postSSML
	}

	if post == nil {
		return errors.New("unknown content type " + contentType.String())
	}

	tsr, err := post(content)
	if err != nil {
		return err
	}

	bts, err := tsr.Bytes()
	if err != nil {
		return err
	}

	_, err = io.Copy(w, bytes.NewReader(bts))

	return err
}

func (s *Synthesizer) WriteText(text string, w io.Writer, _ string) error {
	return s.synthesize(text, tts_integration.Text, w)
}

func (s *Synthesizer) WriteSSML(ssml string, w io.Writer, _ string) error {
	return s.synthesize(ssml, tts_integration.SSML, w)
}

func (s *Synthesizer) VoicesStrings(params ...string) ([]string, error) {
	locale := ""
	if len(params) > 0 {
		locale = params[0]
	}

	vr, err := GetVoices(s.httpClient, locale, s.key)
	if err != nil {
		return nil, err
	}

	voices := make([]string, 0, len(vr.Voices))
	for _, vc := range vr.Voices {
		voices = append(voices, vc.String())
	}

	return voices, nil
}

func (s *Synthesizer) IsSSMLSupported() bool {
	return true
}

func (s *Synthesizer) IsWriterRequired() bool {
	return true
}

func (s *Synthesizer) IsNameRequired() bool {
	return false
}

func (s *Synthesizer) DecorateWithPauses(text string, dur time.Duration) (string, tts_integration.SynthesisInputType) {
	sec := dur.Seconds()
	return fmt.Sprintf("<speak><break time=\"%fs\"/>%s<break time=\"%fs\"/></speak>", sec, text, sec),
		tts_integration.SSML
}
