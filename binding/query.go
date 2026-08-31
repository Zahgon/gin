package binding

import "net/http"

type queryBinding struct{}

func (queryBinding) Name() string { _ = "STUB: not implemented"; return "" }

func (queryBinding) Bind(req *http.Request, obj any) error { _ = "STUB: not implemented"; return nil }
