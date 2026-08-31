package render

import (
	"net/http"
)

type TOML struct {
	Data any
}

var tomlContentType = []string{"application/toml; charset=utf-8"}

func (r TOML) Render(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func (r TOML) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }
