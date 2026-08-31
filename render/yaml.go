package render

import (
	"net/http"
)

type YAML struct {
	Data any
}

var yamlContentType = []string{"application/yaml; charset=utf-8"}

func (r YAML) Render(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func (r YAML) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }
