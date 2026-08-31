package gin

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type consoleColorModeValue int

const (
	autoColor consoleColorModeValue = iota
	disableColor
	forceColor
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

var consoleColorMode = autoColor

type LoggerConfig struct {
	Formatter LogFormatter

	Output io.Writer

	SkipPaths []string

	SkipQueryString bool

	Skip Skipper
}

type Skipper func(c *Context) bool

type LogFormatter func(params LogFormatterParams) string

type LogFormatterParams struct {
	Request *http.Request

	TimeStamp time.Time

	StatusCode int

	Latency time.Duration

	ClientIP string

	Method string

	Path string

	ErrorMessage string

	isTerm bool

	BodySize int

	Keys map[any]any
}

func (p *LogFormatterParams) StatusCodeColor() string { _ = "STUB: not implemented"; return "" }

func (p *LogFormatterParams) LatencyColor() string { _ = "STUB: not implemented"; return "" }

func (p *LogFormatterParams) MethodColor() string { _ = "STUB: not implemented"; return "" }

func (p *LogFormatterParams) ResetColor() string { _ = "STUB: not implemented"; return "" }

func (p *LogFormatterParams) IsOutputColor() bool { _ = "STUB: not implemented"; return false }

var defaultLogFormatter = func(param LogFormatterParams) string {
	var statusColor, methodColor, resetColor, latencyColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
		latencyColor = param.LatencyColor()
	}

	switch {
	case param.Latency > time.Minute:
		param.Latency = param.Latency.Truncate(time.Second * 10)
	case param.Latency > time.Second:
		param.Latency = param.Latency.Truncate(time.Millisecond * 10)
	case param.Latency > time.Millisecond:
		param.Latency = param.Latency.Truncate(time.Microsecond * 10)
	}

	return fmt.Sprintf("[GIN] %v |%s %3d %s|%s %8v %s| %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		latencyColor, param.Latency, resetColor,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

func DisableConsoleColor() { _ = "STUB: not implemented"; return }

func ForceConsoleColor() { _ = "STUB: not implemented"; return }

func ErrorLogger() HandlerFunc { _ = "STUB: not implemented"; return *new(HandlerFunc) }

func ErrorLoggerT(typ ErrorType) HandlerFunc { _ = "STUB: not implemented"; return *new(HandlerFunc) }

func Logger() HandlerFunc { _ = "STUB: not implemented"; return *new(HandlerFunc) }

func LoggerWithFormatter(f LogFormatter) HandlerFunc {
	_ = "STUB: not implemented"
	return *new(HandlerFunc)
}

func LoggerWithWriter(out io.Writer, notlogged ...string) HandlerFunc {
	_ = "STUB: not implemented"
	return *new(HandlerFunc)
}

func LoggerWithConfig(conf LoggerConfig) HandlerFunc {
	_ = "STUB: not implemented"
	return *new(HandlerFunc)
}
