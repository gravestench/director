package file_system

import (
	"io"
	"os"
	"path"
)

func New(basePath string) *Loader {
	result := &Loader{
		basePath: basePath,
	}

	return result
}

type Loader struct {
	basePath string
}

func (l *Loader) Name() string {
	return "FileSystem Loader"
}

func (l *Loader) Exists(p string) bool {
	_, err := os.Stat(path.Join(l.basePath, p))

	return !os.IsNotExist(err)
}

func (l *Loader) Load(p string) (io.ReadSeeker, error) {
	f, err := os.Open(path.Join(l.basePath, p))
	if err != nil {
		return nil, err
	}

	return f, nil
}
