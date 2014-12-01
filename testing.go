package testing

import "archive/zip"
import "bytes"
import "testing"
import "github.com/nowk/assert"

// tZipReader unzips []byte and returns the zip.Reader for the expanded zip
func tZipReader(b []byte) (*zip.Reader, error) {
	r := bytes.NewReader(b)
	z, err := zip.NewReader(r, int64(r.Len()))
	if err != nil {
		return nil, err
	}
	return z, nil
}

// Entries is a simple struct for an entry's name and body contents
type Entries struct {
	Name, Body string
}

// VerifyZip asserts the contents of a zip against an []tTable
func VerifyZip(t *testing.T, b []byte, entries []Entries) error {
	z, err := tZipReader(b)
	if err != nil {
		return err
	}

	for i, entry := range entries {
		f := z.File[i]
		r, err := f.Open()
		if err != nil {
			return err
		}
		defer r.Close()

		b := make([]byte, 32*1024)
		n, err := r.Read(b)
		if err != nil {
			return err
		}

		assert.Equal(t, entry.Name, f.Name)
		assert.Equal(t, entry.Body, string(b[:n]))
	}

	return nil
}
