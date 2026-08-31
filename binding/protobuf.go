package binding

import (
	"net/http"
)

type protobufBinding struct{}

func (protobufBinding) Name() string { _ = "STUB: not implemented"; return "" }

func (b protobufBinding) Bind(req *http.Request, obj any) error {
	_ = "STUB: not implemented"
	return nil
}

func (protobufBinding) BindBody(body []byte, obj any) error { _ = "STUB: not implemented"; return nil }
