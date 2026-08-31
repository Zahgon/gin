package render

import (
	"html/template"
	"net/http"
)

type Delims struct {
	Left string

	Right string
}

type HTMLRender interface {
	Instance(string, any) Render
}

type HTMLProduction struct {
	Template *template.Template
	Delims   Delims
}

type HTMLDebug struct {
	Files      []string
	Glob       string
	FileSystem http.FileSystem
	Patterns   []string
	Delims     Delims
	FuncMap    template.FuncMap
}

type HTML struct {
	Template *template.Template
	Name     string
	Data     any
}

var htmlContentType = []string{"text/html; charset=utf-8"}

func (r HTMLProduction) Instance(name string, data any) Render {
	_ = "STUB: not implemented"
	return *new(Render)
}

func (r HTMLDebug) Instance(name string, data any) Render {
	_ = "STUB: not implemented"
	return *new(Render)
}

func (r HTMLDebug) loadTemplate() *template.Template { _ = "STUB: not implemented"; return nil }

func (r HTML) Render(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func (r HTML) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }
