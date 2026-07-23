//go:build jsoniter

package json

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

const Package = "github.com/json-iterator/go"

func init() {
	API = jsoniterApi{}
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type jsoniterApi struct{}

func (j jsoniterApi) Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (j jsoniterApi) Unmarshal(data []byte, v any) error { _ = "STUB: not implemented"; return nil }

func (j jsoniterApi) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j jsoniterApi) NewEncoder(writer io.Writer) Encoder {
	_ = "STUB: not implemented"
	return *new(Encoder)
}

func (j jsoniterApi) NewDecoder(reader io.Reader) Decoder {
	_ = "STUB: not implemented"
	return *new(Decoder)
}
