package binding

type uriBinding struct{}

func (uriBinding) Name() string { _ = "STUB: not implemented"; return "" }

func (uriBinding) BindUri(m map[string][]string, obj any) error {
	_ = "STUB: not implemented"
	return nil
}
