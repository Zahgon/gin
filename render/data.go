package render

import (
	"net/http"
)

type Data struct {
	ContentType string
	Data        []byte
}

func (r Data) Render(w http.ResponseWriter) (err error) { _ = "STUB: not implemented"; return nil }

func (r Data) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }
