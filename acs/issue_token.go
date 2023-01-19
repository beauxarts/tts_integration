package acs

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func IssueToken(hc *http.Client, region, key string) (string, error) {
	itu := IssueTokenUrl(region)

	issueTokenReq, err := http.NewRequest(http.MethodPost, itu.String(), nil)
	if err != nil {
		return "", err
	}

	issueTokenReq.Header.Add(OcpApimSubscriptionKeyHeader, key)

	resp, err := hc.Do(issueTokenReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return "", errors.New(resp.Status)
	}

	sb := &strings.Builder{}

	_, err = io.Copy(sb, resp.Body)

	return sb.String(), err
}
