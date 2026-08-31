//go:build !nomsgpack

package binding

import (
	"io"
	"net/http"
)

type msgpackBinding struct{}

func (msgpackBinding) Name() string { _ = "STUB: not implemented"; return "" }

func (msgpackBinding) Bind(req *http.Request, obj any) error { _ = "STUB: not implemented"; return nil }

func (msgpackBinding) BindBody(body []byte, obj any) error { _ = "STUB: not implemented"; return nil }

func decodeMsgPack(r io.Reader, obj any) error { _ = "STUB: not implemented"; return nil }
