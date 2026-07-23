//go:build !nomsgpack

package render

import (
	"net/http"
)

var (
	_ Render = MsgPack{}
)

type MsgPack struct {
	Data any
}

var msgpackContentType = []string{"application/msgpack; charset=utf-8"}

func (r MsgPack) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func (r MsgPack) Render(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

func WriteMsgPack(w http.ResponseWriter, obj any) error { _ = "STUB: not implemented"; return nil }
