package gcp

import (
	"bytes"
	"encoding/base64"
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

func NewSynthesizer(hc *http.Client, key string, voiceParams ...string) tts_integration.Synthesizer {
	return &Synthesizer{
		httpClient: hc,
		voice:      NewVoice(voiceParams...),
		key:        key,
	}
}

func (s *Synthesizer) postText(text string) (*TextSynthesizeResponse, error) {
	return TextSynthesize(s.httpClient, text, s.voice, tts_integration.Text, s.key)
}

func (s *Synthesizer) postSSML(ssml string) (*TextSynthesizeResponse, error) {
	return TextSynthesize(s.httpClient, ssml, s.voice, tts_integration.SSML, s.key)
}

func (s *Synthesizer) synthesize(content string, contentType tts_integration.SynthesisInputType, w io.Writer) error {

	var post func(string) (*TextSynthesizeResponse, error)
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
