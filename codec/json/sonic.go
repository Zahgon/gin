//go:build sonic && (linux || windows || darwin)

package json

import (
	"io"

	"github.com/bytedance/sonic"
)

const Package = "github.com/bytedance/sonic"

func init() {
	API = sonicApi{}
}

var json = sonic.ConfigStd

type sonicApi struct{}

func (j sonicApi) Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (j sonicApi) Unmarshal(data []byte, v any) error { _ = "STUB: not implemented"; return nil }

func (j sonicApi) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j sonicApi) NewEncoder(writer io.Writer) Encoder {
	_ = "STUB: not implemented"
	return *new(Encoder)
}

func (j sonicApi) NewDecoder(reader io.Reader) Decoder {
	_ = "STUB: not implemented"
	return *new(Decoder)
}
