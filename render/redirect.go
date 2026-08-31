package render

import (
	"net/http"
)

type Redirect struct {
	Code     int
	Request  *http.Request
	Location string
}

func (r Redirect) Render(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func (r Redirect) WriteContentType(http.ResponseWriter) { _ = "STUB: not implemented"; return }
