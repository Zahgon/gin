package binding

import (
	"net/http"
)

const defaultMemory = 32 << 20

type (
	formBinding          struct{}
	formPostBinding      struct{}
	formMultipartBinding struct{}
)

func (formBinding) Name() string { _ = "STUB: not implemented"; return "" }

func (formBinding) Bind(req *http.Request, obj any) error { _ = "STUB: not implemented"; return nil }

func (formPostBinding) Name() string { _ = "STUB: not implemented"; return "" }

func (formPostBinding) Bind(req *http.Request, obj any) error {
	_ = "STUB: not implemented"
	return nil
}

func (formMultipartBinding) Name() string { _ = "STUB: not implemented"; return "" }

func (formMultipartBinding) Bind(req *http.Request, obj any) error {
	_ = "STUB: not implemented"
	return nil
}
