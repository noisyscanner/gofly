package main

import (
	"bytes"
	"compress/gzip"
)

func zipBytes(bs []byte) ([]byte, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(bs); err != nil {
		return []byte{}, err
	}
	if err := gz.Flush(); err != nil {
		return []byte{}, err
	}
	if err := gz.Close(); err != nil {
		return[]byte{}, err
	}

	return b.Bytes(), nil
}