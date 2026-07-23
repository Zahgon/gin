package gin

import (
	"html/template"
	"net"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin/render"
)

const (
	defaultMultipartMemory = 32 << 20
	escapedColon           = "\\:"
	colon                  = ":"
	backslash              = "\\"
)

var (
	default404Body = []byte("404 page not found")
	default405Body = []byte("405 method not allowed")
)

var defaultPlatform string

var defaultTrustedCIDRs = []*net.IPNet{
	{
		IP:   net.IP{0x0, 0x0, 0x0, 0x0},
		Mask: net.IPMask{0x0, 0x0, 0x0, 0x0},
	},
	{
		IP:   net.IP{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		Mask: net.IPMask{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	},
}

type HandlerFunc func(*Context)

type OptionFunc func(*Engine)

type HandlersChain []HandlerFunc

func (c HandlersChain) Last() HandlerFunc { _ = "STUB: not implemented"; return *new(HandlerFunc) }

type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	HandlerFunc HandlerFunc
}

type RoutesInfo []RouteInfo

const (
	PlatformGoogleAppEngine = "X-Appengine-Remote-Addr"

	PlatformCloudflare = "CF-Connecting-IP"

	PlatformFlyIO = "Fly-Client-IP"
)

type Engine struct {
	RouterGroup

	routeTreesUpdated sync.Once

	RedirectTrailingSlash bool

	RedirectFixedPath bool

	HandleMethodNotAllowed bool

	ForwardedByClientIP bool

	AppEngine bool

	UseRawPath bool

	UseEscapedPath bool

	UnescapePathValues bool

	RemoveExtraSlash bool

	RemoteIPHeaders []string

	TrustedPlatform string

	MaxMultipartMemory int64

	UseH2C bool

	ContextWithFallback bool

	delims           render.Delims
	secureJSONPrefix string
	HTMLRender       render.HTMLRender
	FuncMap          template.FuncMap
	allNoRoute       HandlersChain
	allNoMethod      HandlersChain
	noRoute          HandlersChain
	noMethod         HandlersChain
	pool             sync.Pool
	trees            methodTrees
	maxParams        uint16
	maxSections      uint16
	trustedProxies   []string
	trustedCIDRs     []*net.IPNet
}

var _ IRouter = (*Engine)(nil)

func New(opts ...OptionFunc) *Engine { _ = "STUB: not implemented"; return nil }

func Default(opts ...OptionFunc) *Engine { _ = "STUB: not implemented"; return nil }

func (engine *Engine) Handler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func (engine *Engine) allocateContext(maxParams uint16) *Context {
	_ = "STUB: not implemented"
	return nil
}

func (engine *Engine) Delims(left, right string) *Engine { _ = "STUB: not implemented"; return nil }

func (engine *Engine) SecureJsonPrefix(prefix string) *Engine {
	_ = "STUB: not implemented"
	return nil
}

func (engine *Engine) LoadHTMLGlob(pattern string) { _ = "STUB: not implemented"; return }

func (engine *Engine) LoadHTMLFiles(files ...string) { _ = "STUB: not implemented"; return }

func (engine *Engine) LoadHTMLFS(fs http.FileSystem, patterns ...string) {
	_ = "STUB: not implemented"
	return
}

func (engine *Engine) SetHTMLTemplate(templ *template.Template) { _ = "STUB: not implemented"; return }

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) { _ = "STUB: not implemented"; return }

func (engine *Engine) NoRoute(handlers ...HandlerFunc) { _ = "STUB: not implemented"; return }

func (engine *Engine) NoMethod(handlers ...HandlerFunc) { _ = "STUB: not implemented"; return }

func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (engine *Engine) With(opts ...OptionFunc) *Engine { _ = "STUB: not implemented"; return nil }

func (engine *Engine) rebuild404Handlers() { _ = "STUB: not implemented"; return }

func (engine *Engine) rebuild405Handlers() { _ = "STUB: not implemented"; return }

func (engine *Engine) addRoute(method, path string, handlers HandlersChain) {
	_ = "STUB: not implemented"
	return
}

func (engine *Engine) Routes() (routes RoutesInfo) {
	_ = "STUB: not implemented"
	return *new(RoutesInfo)
}

func iterate(path, method string, routes RoutesInfo, root *node) RoutesInfo {
	_ = "STUB: not implemented"
	return *new(RoutesInfo)
}

func (engine *Engine) prepareTrustedCIDRs() ([]*net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (engine *Engine) SetTrustedProxies(trustedProxies []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (engine *Engine) isUnsafeTrustedProxies() bool { _ = "STUB: not implemented"; return false }

func (engine *Engine) parseTrustedProxies() error { _ = "STUB: not implemented"; return nil }

func (engine *Engine) isTrustedProxy(ip net.IP) bool { _ = "STUB: not implemented"; return false }

func (engine *Engine) validateHeader(header string) (clientIP string, valid bool) {
	_ = "STUB: not implemented"
	return "", false
}

func updateRouteTree(n *node) { _ = "STUB: not implemented"; return }

func (engine *Engine) updateRouteTrees() { _ = "STUB: not implemented"; return }

func parseIP(ip string) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func (engine *Engine) Run(addr ...string) (err error) { _ = "STUB: not implemented"; return nil }

func (engine *Engine) RunTLS(addr, certFile, keyFile string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (engine *Engine) RunUnix(file string) (err error) { _ = "STUB: not implemented"; return nil }

func (engine *Engine) RunFd(fd int) (err error) { _ = "STUB: not implemented"; return nil }

func (engine *Engine) RunQUIC(addr, certFile, keyFile string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (engine *Engine) RunListener(listener net.Listener) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (engine *Engine) HandleContext(c *Context) { _ = "STUB: not implemented"; return }

func (engine *Engine) handleHTTPRequest(c *Context) { _ = "STUB: not implemented"; return }

var mimePlain = []string{MIMEPlain}

func serveError(c *Context, code int, defaultMessage []byte) { _ = "STUB: not implemented"; return }

func redirectTrailingSlash(c *Context) { _ = "STUB: not implemented"; return }

func sanitizePathChars(s string) string { _ = "STUB: not implemented"; return "" }

func redirectFixedPath(c *Context, root *node, trailingSlash bool) bool {
	_ = "STUB: not implemented"
	return false
}

func redirectRequest(c *Context) { _ = "STUB: not implemented"; return }
