package factorio

import (
	"compress/zlib"
	"encoding/base64"
	"io"
)

// Reader is a reader for Factorio blueprint strings. It reads the version byte and returns the uncompressed JSON data.
type Reader struct {
	version byte
	r       io.Reader
	z       io.ReadCloser
}

// NewReader creates a new Reader for the given io.Reader.
func NewReader(r io.Reader) (*Reader, error) {
	version, err := readVersion(r)
	if err != nil {
		return nil, err
	}

	reader, err := zlib.NewReader(base64.NewDecoder(base64.StdEncoding, r))
	if err != nil {
		return nil, err
	}

	return &Reader{
		version: version,
		r:       r,
		z:       reader,
	}, nil
}

func (r *Reader) Close() error {
	return r.z.Close()
}

// Version returns the version byte of the blueprint string.
func (r *Reader) Version() byte {
	return r.version
}

func readVersion(r io.Reader) (byte, error) {
	buf := make([]byte, 1)
	_, err := r.Read(buf)
	if err != nil {
		return 0, err
	}

	return buf[0], nil
}

// Read reads the uncompressed JSON data from the Reader.
func (r *Reader) Read(p []byte) (n int, err error) {
	return r.z.Read(p)
}
