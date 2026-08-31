//go:build nomsgpack

package binding

import "net/http"

const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	MIMEPROTOBUF          = "application/x-protobuf"
	MIMEYAML              = "application/x-yaml"
	MIMEYAML2             = "application/yaml"
	MIMETOML              = "application/toml"
	MIMEBSON              = "application/bson"
)

type Binding interface {
	Name() string
	Bind(*http.Request, any) error
}

type BindingBody interface {
	Binding
	BindBody([]byte, any) error
}

type BindingUri interface {
	Name() string
	BindUri(map[string][]string, any) error
}

type StructValidator interface {
	ValidateStruct(any) error

	Engine() any
}

var Validator StructValidator = &defaultValidator{}

var (
	JSON                      = jsonBinding{}
	XML                       = xmlBinding{}
	Form                      = formBinding{}
	Query                     = queryBinding{}
	FormPost                  = formPostBinding{}
	FormMultipart             = formMultipartBinding{}
	ProtoBuf                  = protobufBinding{}
	YAML                      = yamlBinding{}
	Uri                       = uriBinding{}
	Header                    = headerBinding{}
	TOML                      = tomlBinding{}
	Plain                     = plainBinding{}
	BSON          BindingBody = bsonBinding{}
)

func Default(method, contentType string) Binding { _ = "STUB: not implemented"; return *new(Binding) }

func validate(obj any) error { _ = "STUB: not implemented"; return nil }
