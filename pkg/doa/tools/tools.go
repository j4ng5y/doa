package tools

import (
	"fmt"
	"regexp"
	"strings"
)

type SemVer struct {
	Major string
	Minor string
	Patch string
}

func NewSemVer(v string) (*SemVer, error) {
	r := regexp.MustCompile("(.+?)\\.(.+?)\\.(.+?)")
	if r.MatchString(v) {
		splitz := strings.Split(v, ".")
		return &SemVer{
			Major: splitz[0],
			Minor: splitz[1],
			Patch: splitz[2],
		}, nil
	}
	return nil, fmt.Errorf("%s can not be parsed as SemVer", v)
}

type Tool interface {
	GetLatestVersion(srcURL string) (*SemVer, error)
	Install(srcURL string, version *SemVer) error
	Update(srcURL string, version *SemVer) error
	Remove() error
}
