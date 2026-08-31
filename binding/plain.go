package binding

import (
	"net/http"
)

type plainBinding struct{}

func (plainBinding) Name() string { _ = "STUB: not implemented"; return "" }

func (plainBinding) Bind(req *http.Request, obj any) error { _ = "STUB: not implemented"; return nil }

func (plainBinding) BindBody(body []byte, obj any) error { _ = "STUB: not implemented"; return nil }

func decodePlain(data []byte, obj any) error { _ = "STUB: not implemented"; return nil }
