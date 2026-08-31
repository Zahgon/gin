package render

import (
	"net/http"
)

type String struct {
	Format string
	Data   []any
}

var plainContentType = []string{"text/plain; charset=utf-8"}

func (r String) Render(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func (r String) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func WriteString(w http.ResponseWriter, format string, data []any) (err error) {
	_ = "STUB: not implemented"
	return nil
}
