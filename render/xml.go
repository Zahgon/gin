package render

import (
	"net/http"
)

type XML struct {
	Data any
}

var xmlContentType = []string{"application/xml; charset=utf-8"}

func (r XML) Render(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func (r XML) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }
