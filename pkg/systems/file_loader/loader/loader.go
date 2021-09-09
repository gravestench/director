package loader

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type Loader struct {
	Providers []Provider
}

func (l *Loader) AddProvider(provider Provider) {
	l.Providers = append(l.Providers, provider)
}

func (l *Loader) Load(path string) (io.ReadSeeker, error) {
	if len(path) == 0 {
		return nil, errors.New("blank path provided")
	}

	path = strings.ReplaceAll(path, "\\", "/")

	for providerIdx := range l.Providers {
		if !l.Providers[providerIdx].Exists(path) {
			continue
		}

		return l.Providers[providerIdx].Load(path)
	}

	return nil, fmt.Errorf("file not found: \"%s\"", path)
}
