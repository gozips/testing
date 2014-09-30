package testing

import "archive/zip"
import "bytes"
import "testing"
import "github.com/nowk/assert"

// tZipReader unzips []byte and returns the zip.Reader for the expanded zip
func tZipReader(b []byte) *zip.Reader {
	r := bytes.NewReader(b)
	z, _ := zip.NewReader(r, int64(r.Len()))

	return z
}

// Entries is a simple struct for an entry's name and body contents
type Entries struct {
	Name, Body string
}

// VerifyZip asserts the contents of a zip against an []tTable
func VerifyZip(t *testing.T, b []byte, entries []Entries) {
	z := tZipReader(b)
	for i, entry := range entries {
		f := z.File[i]
		rc, _ := f.Open()
		defer rc.Close()

		b := make([]byte, 32*1024)
		n, _ := rc.Read(b)
		assert.Equal(t, entry.Name, f.Name)
		assert.Equal(t, entry.Body, string(b[:n]))
	}
}
