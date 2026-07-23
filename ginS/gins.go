package ginS

import (
	"html/template"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var engine = sync.OnceValue(func() *gin.Engine {
	return gin.Default()
})

func LoadHTMLGlob(pattern string) { _ = "STUB: not implemented"; return }

func LoadHTMLFiles(files ...string) { _ = "STUB: not implemented"; return }

func LoadHTMLFS(fs http.FileSystem, patterns ...string) { _ = "STUB: not implemented"; return }

func SetHTMLTemplate(templ *template.Template) { _ = "STUB: not implemented"; return }

func NoRoute(handlers ...gin.HandlerFunc) { _ = "STUB: not implemented"; return }

func NoMethod(handlers ...gin.HandlerFunc) { _ = "STUB: not implemented"; return }

func Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	_ = "STUB: not implemented"
	return nil
}

func Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	_ = "STUB: not implemented"
	return *new(gin.IRoutes)
}

func POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	_ = "STUB: not implemented"
	return *new(gin.IRoutes)
}

func GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	_ = "STUB: not implemented"
	return *new(gin.IRoutes)
}

func DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	_ = "STUB: not implemented"
	return *new(gin.IRoutes)
}

func PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	_ = "STUB: not implemented"
	return *new(gin.IRoutes)
}

func PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	_ = "STUB: not implemented"
	return *new(gin.IRoutes)
}

func OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	_ = "STUB: not implemented"
	return *new(gin.IRoutes)
}

func HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	_ = "STUB: not implemented"
	return *new(gin.IRoutes)
}

func Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	_ = "STUB: not implemented"
	return *new(gin.IRoutes)
}

func StaticFile(relativePath, filepath string) gin.IRoutes {
	_ = "STUB: not implemented"
	return *new(gin.IRoutes)
}

func Static(relativePath, root string) gin.IRoutes {
	_ = "STUB: not implemented"
	return *new(gin.IRoutes)
}

func StaticFS(relativePath string, fs http.FileSystem) gin.IRoutes {
	_ = "STUB: not implemented"
	return *new(gin.IRoutes)
}

func Use(middlewares ...gin.HandlerFunc) gin.IRoutes {
	_ = "STUB: not implemented"
	return *new(gin.IRoutes)
}

func Routes() gin.RoutesInfo { _ = "STUB: not implemented"; return *new(gin.RoutesInfo) }

func Run(addr ...string) (err error) { _ = "STUB: not implemented"; return nil }

func RunTLS(addr, certFile, keyFile string) (err error) { _ = "STUB: not implemented"; return nil }

func RunUnix(file string) (err error) { _ = "STUB: not implemented"; return nil }

func RunFd(fd int) (err error) { _ = "STUB: not implemented"; return nil }
