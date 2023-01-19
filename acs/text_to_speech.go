package acs

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	ssmlVersion        = "1.0"
	applicationSsmlXml = "application/ssml+xml"
)

type VoiceParams struct {
	Name   string `xml:"name,attr"`
	Lang   string `xml:"xml:lang,attr,omitempty"`
	Gender string `xml:"xml:gender,attr,omitempty"`
	Text   string `xml:",chardata"`
}

func NewVoiceParams(params ...string) *VoiceParams {
	var voiceName, voiceLang, voiceGender string
	if len(params) > 0 {
		voiceName = params[0]
	}
	if len(params) > 1 {
		voiceLang = params[1]
	}
	if len(params) > 2 {
		voiceGender = params[2]
	}

	return &VoiceParams{
		Name:   voiceName,
		Lang:   voiceLang,
		Gender: voiceGender,
	}
}

func (vp *VoiceParams) NewSpeakData(text string) *SpeakData {
	vp.Text = text
	return &SpeakData{
		Version: ssmlVersion,
		Lang:    vp.Lang,
		Voice:   vp,
	}
}

func (vp *VoiceParams) TextToSpeech(hc *http.Client, text string, audioOutput AudioOutput, region, token string) (io.ReadCloser, error) {
	speakData := vp.NewSpeakData(text)

	data, err := xml.Marshal(speakData)
	if err != nil {
		return nil, err
	}

	return SsmlToSpeech(hc, data, audioOutput, region, token)
}

func (vp *VoiceParams) SsmlContentToSpeech(hc *http.Client, ssml string, audioOutput AudioOutput, region, token string) (io.ReadCloser, error) {
	speakData := vp.NewSpeakData("{ssml}")

	data, err := xml.Marshal(speakData)
	if err != nil {
		return nil, err
	}

	data = []byte(strings.Replace(string(data), "{ssml}", ssml, -1))

	return SsmlToSpeech(hc, data, audioOutput, region, token)
}

type SpeakData struct {
	XMLName xml.Name     `xml:"speak"`
	Version string       `xml:"version,attr"`
	Lang    string       `xml:"xml:lang,attr,omitempty"`
	Voice   *VoiceParams `xml:"voice"`
}

func NewSpeakData(text string, voiceParams ...string) *SpeakData {

	vp := NewVoiceParams(voiceParams...)
	vp.Text = text

	return &SpeakData{
		Version: ssmlVersion,
		Lang:    vp.Lang,
		Voice:   vp,
	}
}

func TextToSpeech(hc *http.Client, text string, audioOutput AudioOutput, region, token string, voiceParams ...string) (io.ReadCloser, error) {

	speakData := NewSpeakData(text, voiceParams...)

	data, err := xml.Marshal(speakData)
	if err != nil {
		return nil, err
	}

	return SsmlToSpeech(hc, data, audioOutput, region, token)
}

func SsmlToSpeech(hc *http.Client, ssml []byte, audioOutput AudioOutput, region, token string) (io.ReadCloser, error) {
	ttsu := TextToSpeechUrl(region)

	fmt.Println(ttsu)

	ttsReq, err := http.NewRequest(http.MethodPost, ttsu.String(), bytes.NewReader(ssml))
	if err != nil {
		return nil, err
	}

	ttsReq.Header.Add("X-Microsoft-OutputFormat", audioOutput.String())
	ttsReq.Header.Add("Content-Type", applicationSsmlXml)
	ttsReq.Header.Add("Authorization", "Bearer "+token)

	resp, err := hc.Do(ttsReq)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	return resp.Body, nil
}
