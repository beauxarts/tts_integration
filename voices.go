package tts_integration

import (
	"encoding/json"
	"net/http"
)

type VoicesResp struct {
	Voices []struct {
		LanguageCodes          []string `json:"languageCodes"`
		Name                   string   `json:"name"`
		SSMLGender             string   `json:"ssmlGender"`
		NaturalSampleRateHertz int      `json:"naturalSampleRateHertz"`
	} `json:"voices"`
}

func GetVoices(hc *http.Client, lc, key string) (*VoicesResp, error) {

	vu := VoicesUrl(lc, key)

	resp, err := hc.Get(vu.String())
	if err != nil {
		return nil, err
	}

	var vs *VoicesResp
	json.NewDecoder(resp.Body).Decode(&vs)

	return vs, err
}
