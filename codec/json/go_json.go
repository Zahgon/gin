//go:build go_json

package json

import (
	"io"
)

const Package = "github.com/goccy/go-json"

func init() {
	API = gojsonApi{}
}

type gojsonApi struct{}

func (j gojsonApi) Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (j gojsonApi) Unmarshal(data []byte, v any) error { _ = "STUB: not implemented"; return nil }

func (j gojsonApi) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j gojsonApi) NewEncoder(writer io.Writer) Encoder {
	_ = "STUB: not implemented"
	return *new(Encoder)
}

func (j gojsonApi) NewDecoder(reader io.Reader) Decoder {
	_ = "STUB: not implemented"
	return *new(Decoder)
}
