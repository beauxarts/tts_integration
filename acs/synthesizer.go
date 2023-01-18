package acs

import (
	"fmt"
	"github.com/beauxarts/tts_integration"
	"io"
	"net/http"
	"strings"
	"time"
)

type Synthesizer struct {
	httpClient  *http.Client
	voice       *VoiceParams
	audioOutput AudioOutput
	region      string
	key         string
}

func NewSynthesizer(hc *http.Client, region, key string, audioOutput AudioOutput, voiceParams ...string) tts_integration.Synthesizer {
	return &Synthesizer{
		httpClient:  hc,
		audioOutput: audioOutput,
		voice:       NewVoiceParams(voiceParams...),
		region:      region,
		key:         key,
	}
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

func (s *Synthesizer) VoicesStrings(params ...string) ([]string, error) {
	locale := ""
	if len(params) > 0 {
		locale = params[0]
	}

	vlr, err := VoicesList(s.httpClient, s.region, s.key)
	if err != nil {
		return nil, err
	}

	voices := make([]string, len(vlr))

	for i, v := range vlr {
		if locale != "" && !strings.Contains(v.Locale, locale) {
			continue
		}
		voices[i] = v.String()
	}

	return voices, nil
}

func (s *Synthesizer) WriteText(text string, w io.Writer, _ string) error {
	rc, err := s.voice.TextToSpeech(s.httpClient, text, s.audioOutput, s.region, s.key)
	if err != nil {
		return err
	}
	defer rc.Close()

	_, err = io.Copy(w, rc)
	return err
}

func (s *Synthesizer) WriteSSML(ssml string, w io.Writer, _ string) error {
	rc, err := s.voice.SsmlSnippetToSpeech(s.httpClient, ssml, s.audioOutput, s.region, s.key)
	if err != nil {
		return err
	}
	defer rc.Close()

	_, err = io.Copy(w, rc)
	return err
}

func (s *Synthesizer) DecorateWithPauses(text string, d time.Duration) (string, tts_integration.SynthesisInputType) {
	sec := d.Seconds()
	return fmt.Sprintf("<break time=\"%fs\"/>%s<break time=\"%fs\"/>", sec, text, sec),
		tts_integration.SSML
}
