package render

import (
	"net/http"
)

type JSON struct {
	Data any
}

type IndentedJSON struct {
	Data any
}

type SecureJSON struct {
	Prefix string
	Data   any
}

type JsonpJSON struct {
	Callback string
	Data     any
}

type AsciiJSON struct {
	Data any
}

type PureJSON struct {
	Data any
}

var (
	jsonContentType      = []string{"application/json; charset=utf-8"}
	jsonpContentType     = []string{"application/javascript; charset=utf-8"}
	jsonASCIIContentType = []string{"application/json"}
)

func (r JSON) Render(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func (r JSON) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func WriteJSON(w http.ResponseWriter, obj any) error { _ = "STUB: not implemented"; return nil }

func (r IndentedJSON) Render(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func (r IndentedJSON) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func (r SecureJSON) Render(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func (r SecureJSON) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func (r JsonpJSON) Render(w http.ResponseWriter) (err error) { _ = "STUB: not implemented"; return nil }

func (r JsonpJSON) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func (r AsciiJSON) Render(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func (r AsciiJSON) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func (r PureJSON) Render(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func (r PureJSON) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }
