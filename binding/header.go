package binding

import (
	"net/http"
	"reflect"
)

type headerBinding struct{}

func (headerBinding) Name() string { _ = "STUB: not implemented"; return "" }

func (headerBinding) Bind(req *http.Request, obj any) error { _ = "STUB: not implemented"; return nil }

func mapHeader(ptr any, h map[string][]string) error { _ = "STUB: not implemented"; return nil }

type headerSource map[string][]string

var _ setter = headerSource(nil)

func (hs headerSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt setOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
