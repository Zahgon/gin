package binding

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

type SliceValidationError []error

func (err SliceValidationError) Error() string { _ = "STUB: not implemented"; return "" }

var _ StructValidator = (*defaultValidator)(nil)

func (v *defaultValidator) ValidateStruct(obj any) error { _ = "STUB: not implemented"; return nil }

func (v *defaultValidator) validateStruct(obj any) error { _ = "STUB: not implemented"; return nil }

func (v *defaultValidator) Engine() any { _ = "STUB: not implemented"; return *new(any) }

func (v *defaultValidator) lazyinit() { _ = "STUB: not implemented"; return }
