package render

import "net/http"

type PDF struct {
	Data []byte
}

var pdfContentType = []string{"application/pdf"}

func (r PDF) Render(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func (r PDF) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }
