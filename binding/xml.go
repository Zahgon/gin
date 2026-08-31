package binding

import (
	"io"
	"net/http"
)

type xmlBinding struct{}

func (xmlBinding) Name() string { _ = "STUB: not implemented"; return "" }

func (xmlBinding) Bind(req *http.Request, obj any) error { _ = "STUB: not implemented"; return nil }

func (xmlBinding) BindBody(body []byte, obj any) error { _ = "STUB: not implemented"; return nil }

func decodeXML(r io.Reader, obj any) error { _ = "STUB: not implemented"; return nil }
