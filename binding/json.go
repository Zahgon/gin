package binding

import (
	"io"
	"net/http"
)

var EnableDecoderUseNumber = false

var EnableDecoderDisallowUnknownFields = false

type jsonBinding struct{}

func (jsonBinding) Name() string { _ = "STUB: not implemented"; return "" }

func (jsonBinding) Bind(req *http.Request, obj any) error { _ = "STUB: not implemented"; return nil }

func (jsonBinding) BindBody(body []byte, obj any) error { _ = "STUB: not implemented"; return nil }

func decodeJSON(r io.Reader, obj any) error { _ = "STUB: not implemented"; return nil }
