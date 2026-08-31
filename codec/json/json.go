//go:build !jsoniter && !go_json && !(sonic && (linux || windows || darwin))

package json

import (
	"io"
)

const Package = "encoding/json"

func init() {
	API = jsonApi{}
}

type jsonApi struct{}

func (j jsonApi) Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (j jsonApi) Unmarshal(data []byte, v any) error { _ = "STUB: not implemented"; return nil }

func (j jsonApi) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j jsonApi) NewEncoder(writer io.Writer) Encoder {
	_ = "STUB: not implemented"
	return *new(Encoder)
}

func (j jsonApi) NewDecoder(reader io.Reader) Decoder {
	_ = "STUB: not implemented"
	return *new(Decoder)
}
