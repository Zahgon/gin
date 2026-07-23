package render

import (
	"net/http"
)

type ProtoBuf struct {
	Data any
}

var protobufContentType = []string{"application/x-protobuf"}

func (r ProtoBuf) Render(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func (r ProtoBuf) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }
