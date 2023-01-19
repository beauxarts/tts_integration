package gcp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type VoiceResp struct {
	LanguageCodes          []string `json:"languageCodes"`
	Name                   string   `json:"name"`
	SSMLGender             string   `json:"ssmlGender"`
	NaturalSampleRateHertz int      `json:"naturalSampleRateHertz"`
}

func (vr *VoiceResp) String() string {
	return fmt.Sprintf("%s;%s;%s;%d",
		vr.Name,
		strings.Join(vr.LanguageCodes, ","),
		vr.SSMLGender,
		vr.NaturalSampleRateHertz)
}

type VoicesResp struct {
	Voices []VoiceResp `json:"voices"`
}

func GetVoices(hc *http.Client, langCode, key string) (*VoicesResp, error) {

	vu := VoicesUrl(langCode)

	voicesReq, err := http.NewRequest(http.MethodGet, vu.String(), nil)
	if err != nil {
		return nil, err
	}

	voicesReq.Header.Add("X-goog-api-key", key)

	resp, err := hc.Do(voicesReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New(resp.Status)
	}

	var vs *VoicesResp
	err = json.NewDecoder(resp.Body).Decode(&vs)

	return vs, err
}
