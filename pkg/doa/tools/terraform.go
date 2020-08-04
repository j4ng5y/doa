package tools

import (
	"net/http"

	"launchpad.net/xmlpath"
)

type Terraform struct{}

func (T *Terraform) GetLatestVersion(srcURL string) (*SemVer, error) {
	resp, err := http.Get(srcURL)
	if err != nil {
		return nil, err
	}

	p := xmlpath.MustCompile("/html/body/ul/li/a")
	root, err := xmlpath.Parse(resp.Body)

	return nil, nil
}
