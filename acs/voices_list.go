package acs

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type VoicesListResponse struct {
	Name                string   `json:"Name"`
	DisplayName         string   `json:"DisplayName"`
	LocalName           string   `json:"LocalName"`
	ShortName           string   `json:"ShortName"`
	Gender              string   `json:"Gender"`
	Locale              string   `json:"Locale"`
	LocaleName          string   `json:"LocaleName"`
	SampleRateHertz     string   `json:"SampleRateHertz"`
	VoiceType           string   `json:"VoiceType"`
	Status              string   `json:"Status"`
	WordsPerMinute      string   `json:"WordsPerMinute,omitempty"`
	StyleList           []string `json:"StyleList,omitempty"`
	SecondaryLocaleList []string `json:"SecondaryLocaleList,omitempty"`
	RolePlayList        []string `json:"RolePlayList,omitempty"`
}

func (vlr *VoicesListResponse) String() string {
	return strings.Join([]string{vlr.ShortName, vlr.Locale, vlr.Gender}, ";")
}

func VoicesList(hc *http.Client, region, token string) ([]*VoicesListResponse, error) {
	vlu := VoicesListUrl(region)

	voicesListReq, err := http.NewRequest(http.MethodGet, vlu.String(), nil)
	if err != nil {
		return nil, err
	}

	voicesListReq.Header.Add("Authorization", "Bearer "+token)

	resp, err := hc.Do(voicesListReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	var vlr []*VoicesListResponse
	err = json.NewDecoder(resp.Body).Decode(&vlr)

	return vlr, err
}
