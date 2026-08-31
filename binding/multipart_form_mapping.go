package binding

import (
	"errors"
	"mime/multipart"
	"net/http"
	"reflect"
)

type multipartRequest http.Request

var _ setter = (*multipartRequest)(nil)

var (
	ErrMultiFileHeader = errors.New("unsupported field type for multipart.FileHeader")

	ErrMultiFileHeaderLenInvalid = errors.New("unsupported len of array for []*multipart.FileHeader")
)

func (r *multipartRequest) TrySet(value reflect.Value, field reflect.StructField, key string, opt setOptions) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func setByMultipartFormFile(value reflect.Value, field reflect.StructField, files []*multipart.FileHeader) (isSet bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func setArrayOfMultipartFormFiles(value reflect.Value, field reflect.StructField, files []*multipart.FileHeader) (isSet bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}
