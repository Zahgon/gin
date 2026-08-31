package render

import (
	"net/http"
)

type BSON struct {
	Data any
}

var bsonContentType = []string{"application/bson"}

func (r BSON) Render(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func (r BSON) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }
