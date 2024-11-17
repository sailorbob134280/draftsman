package factorio

import (
	"bytes"
	"testing"
)

func TestWriterWrite(t *testing.T) {
	// The actual output doesn't appear to be super predictable, but it does work in-game, so we'll just check for errors
	cases := []struct {
		name    string
		json    string
		version byte
		wantErr error
	}{
		{
			name:    "good JSON",
			json:    `{"blueprint":{"icons":[{"signal":{"name":"assembling-machine-3"},"index":1}],"entities":[{"entity_number":1,"name":"assembling-machine-3","position":{"x":-100.5,"y":-222.5}},{"entity_number":2,"name":"bulk-inserter","position":{"x":-101.5,"y":-220.5},"direction":8},{"entity_number":3,"name":"bulk-inserter","position":{"x":-99.5,"y":-220.5}},{"entity_number":4,"name":"requester-chest","position":{"x":-101.5,"y":-219.5},"request_filters":{"sections":[{"index":1}]}},{"entity_number":5,"name":"passive-provider-chest","position":{"x":-99.5,"y":-219.5}}],"item":"blueprint","version":562949954404356}}`,
			version: 48,
			wantErr: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var buf bytes.Buffer
			w := NewWriter(&buf, c.version)
			defer w.Close()

			_, err := w.Write([]byte(c.json))
			if err != c.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, c.wantErr)
			}

			if buf.Bytes()[0] != c.version {
				t.Errorf("Write() version = %v, want %v", buf.Bytes()[0], c.version)
			}
		})
	}
}
