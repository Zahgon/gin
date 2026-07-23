package render

import (
	"io"
	"net/http"
)

type Reader struct {
	ContentType   string
	ContentLength int64
	Reader        io.Reader
	Headers       map[string]string
}

func (r Reader) Render(w http.ResponseWriter) (err error) { _ = "STUB: not implemented"; return nil }

func (r Reader) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func (r Reader) writeHeaders(w http.ResponseWriter) { _ = "STUB: not implemented"; return }
