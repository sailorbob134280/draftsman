package factorio

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"io"
	"strings"
	"testing"
)

func TestNewReader(t *testing.T) {
	cases := []struct {
		name        string
		testString  string
		wantVersion byte
		wantErr     error
	}{
		{
			name:        "valid string",
			testString:  "0eNqNkttugzAMht/F10lVDqkWXmWaKg5eaw0MSwCtQnn3GZjopqFuV0kc/5//2JmgqAfsHHEP2QRUtuwhe57A04Xzeo5x3iBkkHuPTVETX3STl1di1AkEBcQVfkAWhRcFyD31hCthOdzOPDQFOklQD0kKutaLuOW5pgB1dDwejIKbbOM4PpgQ1C9ovEGLoX7TxB5dLxd7tOhOE7A4r8hhueY87bCTf7Ot/YnegaUbzOH7gF5AurzK+ofVyC5Wv0TnV6pF6edMv3pfe30fwl5xsxXvpPc0ou5cO1L1wMO3Jy0W5ulSj83cje3DKBjFzKIyp9im1po0PaaJOYXwCXGDw/Y=",
			wantVersion: 48,
			wantErr:     nil,
		},
		{
			name:        "empty string",
			testString:  "",
			wantVersion: 0,
			wantErr:     io.EOF,
		},
		{
			name:        "bad header",
			testString:  "0ethisheaderisbrooooookenAMht/F10lVDqkWXmWaKg5eaw0MSwCt",
			wantVersion: 48,
			wantErr:     zlib.ErrHeader,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r, err := NewReader(strings.NewReader(c.testString))
			if err != c.wantErr {
				t.Errorf("NewReader() error = %v, wantErr %v", err, c.wantErr)
			}

			if r != nil && r.Version() != c.wantVersion {
				t.Errorf("NewReader() version = %v, want %v", r.Version(), c.wantVersion)
			}
		})
	}
}

func TestWriterRead(t *testing.T) {
	cases := []struct {
		name        string
		testString  string
		want        string
		wantMakeErr error
		wantReadErr error
	}{
		{
			name:        "valid string",
			testString:  "0eNqNkttugzAMht/F10lVDqkWXmWaKg5eaw0MSwCtQnn3GZjopqFuV0kc/5//2JmgqAfsHHEP2QRUtuwhe57A04Xzeo5x3iBkkHuPTVETX3STl1di1AkEBcQVfkAWhRcFyD31hCthOdzOPDQFOklQD0kKutaLuOW5pgB1dDwejIKbbOM4PpgQ1C9ovEGLoX7TxB5dLxd7tOhOE7A4r8hhueY87bCTf7Ot/YnegaUbzOH7gF5AurzK+ofVyC5Wv0TnV6pF6edMv3pfe30fwl5xsxXvpPc0ou5cO1L1wMO3Jy0W5ulSj83cje3DKBjFzKIyp9im1po0PaaJOYXwCXGDw/Y=",
			want:        `{"blueprint":{"icons":[{"signal":{"name":"assembling-machine-3"},"index":1}],"entities":[{"entity_number":1,"name":"assembling-machine-3","position":{"x":-100.5,"y":-222.5}},{"entity_number":2,"name":"bulk-inserter","position":{"x":-101.5,"y":-220.5},"direction":8},{"entity_number":3,"name":"bulk-inserter","position":{"x":-99.5,"y":-220.5}},{"entity_number":4,"name":"requester-chest","position":{"x":-101.5,"y":-219.5},"request_filters":{"sections":[{"index":1}]}},{"entity_number":5,"name":"passive-provider-chest","position":{"x":-99.5,"y":-219.5}}],"item":"blueprint","version":562949954404356}}`,
			wantMakeErr: nil,
			wantReadErr: nil,
		},
		{
			name:        "empty except version",
			testString:  "0e",
			want:        "",
			wantMakeErr: io.ErrUnexpectedEOF,
			wantReadErr: nil,
		},
		{
			name:        "bad file",
			testString:  "0eNqNkttugzAMht/F10lVDqkWXmWaKg5eaw0MSwCtQnn3GZjopqFuV0kc/5//2JmgqAfsHHEP2QRUtuwhe57A04Xzeo5aLuWowLookICorruptedItpgQ1C9ovEGLoX7TxB5dLxd7tOhOE7A4r8hhueY87bCTf7Ot/YnegaUbzOH7gF5AurzK+ofVyC5Wv0TnV6pF6edMv3pfe30fwl5xsxXvpPc0ou5cO1L1wMO3Jy0W5ulSj83cje3DKBjFzKIyp9im1po0PaaJOYXwCXGDw/Y=",
			want:        "",
			wantMakeErr: nil,
			wantReadErr: flate.CorruptInputError(76),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r, err := NewReader(strings.NewReader(c.testString))
			// If there was an error, we don't want to continue reading, but we want to make sure it's the error we expected.
			if err != nil {
				if err != c.wantMakeErr {
					t.Errorf("NewReader() error = %v, expected %v", err, c.wantMakeErr)
				}
				return
			}

			var buf bytes.Buffer
			_, err = io.Copy(&buf, r)
			if err != c.wantReadErr {
				t.Errorf("io.Copy() error = %v, wantErr %v", err, c.wantReadErr)
			}
			if err == nil && buf.String() != c.want {
				t.Errorf("Read() result = %v, want %v", buf.String(), c.want)
			}
		})
	}
}
