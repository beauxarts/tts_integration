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
	token       string
	tokenIssued time.Time
}

func NewSynthesizer(hc *http.Client,
	region, key string,
	audioOutput AudioOutput,
	voiceParams ...string) (tts_integration.Synthesizer, error) {

	s := &Synthesizer{
		httpClient:  hc,
		audioOutput: audioOutput,
		voice:       NewVoiceParams(voiceParams...),
		key:         key,
		region:      region,
	}

	err := s.RefreshToken()

	return s, err
}

func (s *Synthesizer) RefreshToken() error {
	if time.Now().Sub(s.tokenIssued).Minutes() >= 10 {
		token, err := IssueToken(s.httpClient, s.region, s.key)
		if err != nil {
			return err
		}
		s.token = token
		s.tokenIssued = time.Now()
	}
	return nil
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

	if err := s.RefreshToken(); err != nil {
		return nil, err
	}

	vlr, err := VoicesList(s.httpClient, s.region, s.token)
	if err != nil {
		return nil, err
	}

	voices := make([]string, 0, len(vlr))

	for _, v := range vlr {
		if locale != "" && !strings.Contains(v.Locale, locale) {
			continue
		}
		voices = append(voices, v.String())
	}

	return voices, nil
}

func (s *Synthesizer) WriteText(text string, w io.Writer, _ string) error {
	if err := s.RefreshToken(); err != nil {
		return err
	}

	rc, err := s.voice.TextToSpeech(s.httpClient, text, s.audioOutput, s.region, s.token)
	if err != nil {
		return err
	}
	defer rc.Close()

	_, err = io.Copy(w, rc)
	return err
}

func (s *Synthesizer) WriteSSML(ssml string, w io.Writer, _ string) error {
	if err := s.RefreshToken(); err != nil {
		return err
	}

	rc, err := s.voice.SsmlContentToSpeech(s.httpClient, ssml, s.audioOutput, s.region, s.token)
	if err != nil {
		return err
	}
	defer rc.Close()

	_, err = io.Copy(w, rc)
	return err
}

func (s *Synthesizer) DecorateWithPauses(text string, d time.Duration) (string, tts_integration.SynthesisInputType) {
	sec := int(d.Seconds())
	return fmt.Sprintf("<break time=\"%ds\"/>%s<break time=\"%ds\"/>", sec, text, sec),
		tts_integration.SSML
}
