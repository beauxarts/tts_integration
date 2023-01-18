package acs

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
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

type SpeakData struct {
	XMLName xml.Name     `xml:"speak"`
	Version string       `xml:"version,attr"`
	Lang    string       `xml:"xml:lang,attr,omitempty"`
	Voice   *VoiceParams `xml:"voice"`
}

func NewSpeakData(text string, voiceParams ...string) *SpeakData {
	var voiceName, voiceLang, voiceGender string
	if len(voiceParams) > 0 {
		voiceName = voiceParams[0]
	}
	if len(voiceParams) > 1 {
		voiceLang = voiceParams[1]
	}
	if len(voiceParams) > 2 {
		voiceGender = voiceParams[2]
	}

	return &SpeakData{
		Version: ssmlVersion,
		Lang:    voiceLang,
		Voice: &VoiceParams{
			Name:   voiceName,
			Lang:   voiceLang,
			Gender: voiceGender,
			Text:   text,
		},
	}
}

func TextToSpeech(hc *http.Client, text string, audioOutput AudioOutput, region, key string, voiceParams ...string) (io.ReadCloser, error) {

	speakData := NewSpeakData(text, voiceParams...)

	data, err := xml.Marshal(speakData)
	if err != nil {
		return nil, err
	}

	ttsu := TextToSpeechUrl(region)

	ttsReq, err := http.NewRequest(http.MethodPost, ttsu.String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	ttsReq.Header.Add("X-Microsoft-OutputFormat", audioOutput.String())
	ttsReq.Header.Add("Content-Type", applicationSsmlXml)
	ttsReq.Header.Add("Ocp-Apim-Subscription-Key", key)

	resp, err := hc.Do(ttsReq)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	return resp.Body, nil
}
