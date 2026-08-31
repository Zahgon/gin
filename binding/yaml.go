package binding

import (
	"io"
	"net/http"
)

type yamlBinding struct{}

func (yamlBinding) Name() string { _ = "STUB: not implemented"; return "" }

func (yamlBinding) Bind(req *http.Request, obj any) error { _ = "STUB: not implemented"; return nil }

func (yamlBinding) BindBody(body []byte, obj any) error { _ = "STUB: not implemented"; return nil }

func decodeYAML(r io.Reader, obj any) error { _ = "STUB: not implemented"; return nil }
