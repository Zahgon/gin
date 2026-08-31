package binding

import (
	"io"
	"net/http"
)

type tomlBinding struct{}

func (tomlBinding) Name() string { _ = "STUB: not implemented"; return "" }

func (tomlBinding) Bind(req *http.Request, obj any) error { _ = "STUB: not implemented"; return nil }

func (tomlBinding) BindBody(body []byte, obj any) error { _ = "STUB: not implemented"; return nil }

func decodeToml(r io.Reader, obj any) error { _ = "STUB: not implemented"; return nil }
