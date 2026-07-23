package gin

import (
	"io"
	"net/http"
	"time"
)

const (
	dunno     = "???"
	stackSkip = 3
)

type RecoveryFunc func(c *Context, err any)

func Recovery() HandlerFunc { _ = "STUB: not implemented"; return *new(HandlerFunc) }

func CustomRecovery(handle RecoveryFunc) HandlerFunc {
	_ = "STUB: not implemented"
	return *new(HandlerFunc)
}

func RecoveryWithWriter(out io.Writer, recovery ...RecoveryFunc) HandlerFunc {
	_ = "STUB: not implemented"
	return *new(HandlerFunc)
}

func CustomRecoveryWithWriter(out io.Writer, handle RecoveryFunc) HandlerFunc {
	_ = "STUB: not implemented"
	return *new(HandlerFunc)
}

//nolint: errcheck

func secureRequestDump(r *http.Request) string { _ = "STUB: not implemented"; return "" }

func defaultHandleRecovery(c *Context, err any) { _ = "STUB: not implemented"; return }

//nolint: errcheck

func stack(skip int) []byte { _ = "STUB: not implemented"; return nil }

func readNthLine(file string, n int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func function(pc uintptr) string { _ = "STUB: not implemented"; return "" }

func timeFormat(t time.Time) string { _ = "STUB: not implemented"; return "" }
