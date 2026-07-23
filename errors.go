package gin

type ErrorType uint64

const (
	ErrorTypeBind ErrorType = 1 << 63

	ErrorTypeRender ErrorType = 1 << 62

	ErrorTypePrivate ErrorType = 1 << 0

	ErrorTypePublic ErrorType = 1 << 1

	ErrorTypeAny ErrorType = 1<<64 - 1
)

type Error struct {
	Err  error
	Type ErrorType
	Meta any
}

type errorMsgs []*Error

var _ error = (*Error)(nil)

func (msg *Error) SetType(flags ErrorType) *Error { _ = "STUB: not implemented"; return nil }

func (msg *Error) SetMeta(data any) *Error { _ = "STUB: not implemented"; return nil }

func (msg *Error) JSON() any { _ = "STUB: not implemented"; return *new(any) }

func (msg *Error) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (msg Error) Error() string { _ = "STUB: not implemented"; return "" }

func (msg *Error) IsType(flags ErrorType) bool { _ = "STUB: not implemented"; return false }

func (msg Error) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (a errorMsgs) ByType(typ ErrorType) errorMsgs {
	_ = "STUB: not implemented"
	return *new(errorMsgs)
}

func (a errorMsgs) Last() *Error { _ = "STUB: not implemented"; return nil }

func (a errorMsgs) Errors() []string { _ = "STUB: not implemented"; return nil }

func (a errorMsgs) JSON() any { _ = "STUB: not implemented"; return *new(any) }

func (a errorMsgs) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (a errorMsgs) String() string { _ = "STUB: not implemented"; return "" }
