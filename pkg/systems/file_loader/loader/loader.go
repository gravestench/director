package loader

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type Loader struct {
	providers           []Provider
}

func (l *Loader) AddProvider(provider Provider) {
	l.providers = append(l.providers, provider)
}

func (l *Loader) Load(path string) (io.ReadSeeker, error) {
	if len(path) == 0 {
		return nil, errors.New("blank path provided")
	}

	path = strings.ReplaceAll(path, "\\", "/")

	for providerIdx := range l.providers {
		if !l.providers[providerIdx].Exists(path) {
			continue
		}

		return l.providers[providerIdx].Load(path)
	}

	return nil, fmt.Errorf("file not found: \"%s\"", path)
}
