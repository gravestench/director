package loader

import "io"

type Provider interface {
	Name() string
	Exists(path string) bool
	Load(path string) (io.ReadSeeker, error)
}

