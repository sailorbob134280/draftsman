package factorio

import (
	"compress/zlib"
	"encoding/base64"
	"io"
)

// Writer is a writer for Factorio blueprint strings. It accepts uncompressed JSON data and writes the version byte and the compressed data.
type Writer struct {
	version      byte
	w            io.Writer
	e            io.WriteCloser
	z            io.WriteCloser
	wroteVersion bool
}

// NewWriter creates a new Writer for the given io.Writer.
func NewWriter(w io.Writer, version byte) *Writer {
	e := base64.NewEncoder(base64.StdEncoding, w)
	return &Writer{
		version: version,
		w:       w,
		e:       e,
		z:       zlib.NewWriter(e),
	}
}

func (w *Writer) writeVersion() error {
	_, err := w.w.Write([]byte{w.version})
	if err != nil {
		return err
	}

	w.wroteVersion = true
	return nil
}

// Write writes the uncompressed JSON data to the Writer.
func (w *Writer) Write(p []byte) (n int, err error) {
	if !w.wroteVersion {
		err := w.writeVersion()
		if err != nil {
			return 0, err
		}
	}

	e := base64.NewEncoder(base64.StdEncoding, w.w)
	defer e.Close()

	z := zlib.NewWriter(e)
	defer z.Close()

	return z.Write(p)
}

// Close closes the Writer.
func (w *Writer) Close() error {
	if !w.wroteVersion {
		err := w.writeVersion()
		if err != nil {
			return err
		}
	}

	return nil
}
