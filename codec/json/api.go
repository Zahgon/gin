package json

import "io"

var API Core

type Core interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
	MarshalIndent(v any, prefix, indent string) ([]byte, error)
	NewEncoder(writer io.Writer) Encoder
	NewDecoder(reader io.Reader) Decoder
}

type Encoder interface {
	SetEscapeHTML(on bool)

	Encode(v any) error
}

type Decoder interface {
	UseNumber()

	DisallowUnknownFields()

	Decode(v any) error
}
