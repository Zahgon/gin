package gin

import (
	"html/template"
	"runtime"
)

const ginSupportMinGoVer = 25

var runtimeVersion = runtime.Version()

func IsDebugging() bool { _ = "STUB: not implemented"; return false }

var DebugPrintRouteFunc func(httpMethod, absolutePath, handlerName string, nuHandlers int)

var DebugPrintFunc func(format string, values ...any)

func debugPrintRoute(httpMethod, absolutePath string, handlers HandlersChain) {
	_ = "STUB: not implemented"
	return
}

func debugPrintLoadTemplate(tmpl *template.Template) { _ = "STUB: not implemented"; return }

func debugPrint(format string, values ...any) { _ = "STUB: not implemented"; return }

func getMinVer(v string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func debugPrintWARNINGDefault() { _ = "STUB: not implemented"; return }

func debugPrintWARNINGNew() { _ = "STUB: not implemented"; return }

func debugPrintWARNINGSetHTMLTemplate() { _ = "STUB: not implemented"; return }

func debugPrintError(err error) { _ = "STUB: not implemented"; return }
