package binding

import (
	"errors"
	"reflect"
)

var (
	errUnknownType = errors.New("unknown type")

	ErrConvertMapStringSlice = errors.New("can not convert to map slices of strings")

	ErrConvertToMapString = errors.New("can not convert to map of strings")
)

func mapURI(ptr any, m map[string][]string) error { _ = "STUB: not implemented"; return nil }

func mapForm(ptr any, form map[string][]string) error { _ = "STUB: not implemented"; return nil }

func MapFormWithTag(ptr any, form map[string][]string, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

var emptyField = reflect.StructField{}

func mapFormByTag(ptr any, form map[string][]string, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

type setter interface {
	TrySet(value reflect.Value, field reflect.StructField, key string, opt setOptions) (isSet bool, err error)
}

type formSource map[string][]string

var _ setter = formSource(nil)

func (form formSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt setOptions) (isSet bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func mappingByPtr(ptr any, setter setter, tag string) error { _ = "STUB: not implemented"; return nil }

func mapping(value reflect.Value, field reflect.StructField, setter setter, tag string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type setOptions struct {
	isDefaultExists bool
	defaultValue    string

	parser string
}

func tryToSetValue(value reflect.Value, field reflect.StructField, setter setter, tag string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type BindUnmarshaler interface {
	UnmarshalParam(param string) error
}

func trySetCustom(val string, value reflect.Value) (isSet bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func trySetUsingParser(val string, value reflect.Value, parser string) (isSet bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func trySplit(vs []string, field reflect.StructField) (newVs []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setByForm(value reflect.Value, field reflect.StructField, form map[string][]string, tagValue string, opt setOptions) (isSet bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func setWithProperType(val string, value reflect.Value, field reflect.StructField, opt setOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func setIntField(val string, bitSize int, field reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setUintField(val string, bitSize int, field reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setBoolField(val string, field reflect.Value) error { _ = "STUB: not implemented"; return nil }

func setFloatField(val string, bitSize int, field reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setTimeField(val string, structField reflect.StructField, value reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func setArray(vals []string, value reflect.Value, field reflect.StructField, opt setOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func setSlice(vals []string, value reflect.Value, field reflect.StructField, opt setOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func setTimeDuration(val string, value reflect.Value) error { _ = "STUB: not implemented"; return nil }

func head(str, sep string) (head string, tail string) { _ = "STUB: not implemented"; return "", "" }

func setFormMap(ptr any, form map[string][]string) error { _ = "STUB: not implemented"; return nil }
