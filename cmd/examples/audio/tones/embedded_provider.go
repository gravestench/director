package main

import (
	"bytes"
	"embed"
	"io"
)

const (
	bytesToRead = 16
)

type provider struct {
	data *embed.FS
}

func (p *provider) Name() string {
	return "Embedded Filesystem"
}

func (p *provider) Exists(path string) bool {
	f, err := p.data.Open(path)
	if f != nil {
		_ = f.Close()
	}

	return err == nil
}

func (p *provider) Load(path string) (io.ReadSeeker, error) {
	f, err := p.data.Open(path)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, bytesToRead)
	res := make([]byte, 0)

	for {
		numRead, readErr := f.Read(buf)

		res = append(res, buf[:numRead]...)

		if readErr != nil {
			break
		}
	}

	return bytes.NewReader(res), nil
}
