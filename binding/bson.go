package binding

import (
	"net/http"
)

type bsonBinding struct{}

func (bsonBinding) Name() string { _ = "STUB: not implemented"; return "" }

func (b bsonBinding) Bind(req *http.Request, obj any) error { _ = "STUB: not implemented"; return nil }

func (bsonBinding) BindBody(body []byte, obj any) error { _ = "STUB: not implemented"; return nil }
