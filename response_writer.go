package gin

import (
	"bufio"
	"errors"
	"net"
	"net/http"
)

const (
	noWritten     = -1
	defaultStatus = http.StatusOK
)

var errHijackAlreadyWritten = errors.New("gin: response body already written")

type ResponseWriter interface {
	http.ResponseWriter
	http.Hijacker
	http.Flusher
	http.CloseNotifier

	Status() int

	Size() int

	WriteString(string) (int, error)

	Written() bool

	WriteHeaderNow()

	Pusher() http.Pusher
}

type responseWriter struct {
	http.ResponseWriter
	size   int
	status int
}

var _ ResponseWriter = (*responseWriter)(nil)

func (w *responseWriter) Unwrap() http.ResponseWriter {
	_ = "STUB: not implemented"
	return *new(http.ResponseWriter)
}

func (w *responseWriter) reset(writer http.ResponseWriter) { _ = "STUB: not implemented"; return }

func (w *responseWriter) WriteHeader(code int) { _ = "STUB: not implemented"; return }

func (w *responseWriter) WriteHeaderNow() { _ = "STUB: not implemented"; return }

func (w *responseWriter) Write(data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *responseWriter) WriteString(s string) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *responseWriter) Status() int { _ = "STUB: not implemented"; return 0 }

func (w *responseWriter) Size() int { _ = "STUB: not implemented"; return 0 }

func (w *responseWriter) Written() bool { _ = "STUB: not implemented"; return false }

func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

func (w *responseWriter) CloseNotify() <-chan bool { _ = "STUB: not implemented"; return nil }

func (w *responseWriter) Flush() { _ = "STUB: not implemented"; return }

func (w *responseWriter) Pusher() (pusher http.Pusher) {
	_ = "STUB: not implemented"
	return *new(http.Pusher)
}
