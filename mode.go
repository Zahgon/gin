package gin

import (
	"io"
	"os"
	"sync/atomic"
)

const EnvGinMode = "GIN_MODE"

const (
	DebugMode = "debug"

	ReleaseMode = "release"

	TestMode = "test"
)

const (
	debugCode = iota
	releaseCode
	testCode
)

var DefaultWriter io.Writer = os.Stdout

var DefaultErrorWriter io.Writer = os.Stderr

var (
	ginMode  int32 = debugCode
	modeName atomic.Value
)

func init() {
	mode := os.Getenv(EnvGinMode)
	SetMode(mode)
}

func SetMode(value string) { _ = "STUB: not implemented"; return }

func DisableBindValidation() { _ = "STUB: not implemented"; return }

func EnableJsonDecoderUseNumber() { _ = "STUB: not implemented"; return }

func EnableJsonDecoderDisallowUnknownFields() { _ = "STUB: not implemented"; return }

func Mode() string { _ = "STUB: not implemented"; return "" }
