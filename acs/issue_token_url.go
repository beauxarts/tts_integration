package acs

import (
	"net/url"
	"strings"
)

func IssueTokenUrl(region string) *url.URL {
	return &url.URL{
		Scheme: httpsScheme,
		Host:   strings.Replace(regionalApiCognitiveMicrosoftHostTemplate, "{region}", region, -1),
		Path:   issueTokenPath,
	}
}
