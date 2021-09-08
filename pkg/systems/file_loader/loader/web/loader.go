package web

import (
	"bytes"
	"github.com/gravestench/director/pkg/common/cache"
	"io"
	"net/http"
	"net/url"
)

type Loader struct {
	cache *cache.Cache
}

const (
	imageBudget = 100
)

func New() *Loader {
	return &Loader{
		cache: cache.New(imageBudget),
	}
}

func (l *Loader) Name() string {
	return "HTTP Web Loader"
}

func (l *Loader) Exists(p string) bool {
	if _, err := url.Parse(p); err != nil {
		return false
	}

	resp, err := http.Get(p)
	if err != nil {
		return false
	}

	if resp.StatusCode != 200 {
		return false
	}

	return true
}

func (l *Loader) Load(p string) (io.ReadSeeker, error) {
	if entry, found := l.cache.Retrieve(p); found {
		return entry.(io.ReadSeeker), nil
	}

	resp, err := http.Get(p)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)
	 _ = l.cache.Insert(p, reader, 1)

	return reader, nil
}