package gin

import (
	"net/http"
	"regexp"
)

var (
	regEnLetter = regexp.MustCompile("^[A-Z]+$")

	anyMethods = []string{
		http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch,
		http.MethodHead, http.MethodOptions, http.MethodDelete, http.MethodConnect,
		http.MethodTrace,
	}
)

type IRouter interface {
	IRoutes
	Group(string, ...HandlerFunc) *RouterGroup
}

type IRoutes interface {
	Use(...HandlerFunc) IRoutes

	Handle(string, string, ...HandlerFunc) IRoutes
	Any(string, ...HandlerFunc) IRoutes
	GET(string, ...HandlerFunc) IRoutes
	POST(string, ...HandlerFunc) IRoutes
	DELETE(string, ...HandlerFunc) IRoutes
	PATCH(string, ...HandlerFunc) IRoutes
	PUT(string, ...HandlerFunc) IRoutes
	OPTIONS(string, ...HandlerFunc) IRoutes
	HEAD(string, ...HandlerFunc) IRoutes
	Match([]string, string, ...HandlerFunc) IRoutes

	StaticFile(string, string) IRoutes
	StaticFileFS(string, string, http.FileSystem) IRoutes
	Static(string, string) IRoutes
	StaticFS(string, http.FileSystem) IRoutes
}

type RouterGroup struct {
	Handlers HandlersChain
	basePath string
	engine   *Engine
	root     bool
}

var _ IRouter = (*RouterGroup)(nil)

func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	_ = "STUB: not implemented"
	return nil
}

func (group *RouterGroup) BasePath() string { _ = "STUB: not implemented"; return "" }

func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) Any(relativePath string, handlers ...HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) Match(methods []string, relativePath string, handlers ...HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) StaticFile(relativePath, filepath string) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) StaticFileFS(relativePath, filepath string, fs http.FileSystem) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) staticFileHandler(relativePath string, handler HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) Static(relativePath, root string) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) StaticFS(relativePath string, fs http.FileSystem) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	_ = "STUB: not implemented"
	return *new(HandlerFunc)
}

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	_ = "STUB: not implemented"
	return *new(HandlersChain)
}

func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
	_ = "STUB: not implemented"
	return ""
}

func (group *RouterGroup) returnObj() IRoutes { _ = "STUB: not implemented"; return *new(IRoutes) }
