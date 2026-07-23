package gin

import (
	"io"
	"io/fs"
	"math"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin/render"
)

const (
	MIMEJSON              = binding.MIMEJSON
	MIMEHTML              = binding.MIMEHTML
	MIMEXML               = binding.MIMEXML
	MIMEXML2              = binding.MIMEXML2
	MIMEPlain             = binding.MIMEPlain
	MIMEPOSTForm          = binding.MIMEPOSTForm
	MIMEMultipartPOSTForm = binding.MIMEMultipartPOSTForm
	MIMEYAML              = binding.MIMEYAML
	MIMEYAML2             = binding.MIMEYAML2
	MIMETOML              = binding.MIMETOML
	MIMEPROTOBUF          = binding.MIMEPROTOBUF
	MIMEBSON              = binding.MIMEBSON
)

const BodyBytesKey = "_gin-gonic/gin/bodybyteskey"

const ContextKey = "_gin-gonic/gin/contextkey"

type ContextKeyType int

const ContextRequestKey ContextKeyType = 0

const abortIndex int8 = math.MaxInt8 >> 1

type Context struct {
	writermem responseWriter
	Request   *http.Request
	Writer    ResponseWriter

	Params   Params
	handlers HandlersChain
	index    int8
	fullPath string

	engine       *Engine
	params       *Params
	skippedNodes *[]skippedNode

	mu sync.RWMutex

	Keys map[any]any

	Errors errorMsgs

	Accepted []string

	queryCache url.Values

	formCache url.Values

	sameSite http.SameSite
}

func (c *Context) reset() { _ = "STUB: not implemented"; return }

func (c *Context) Copy() *Context { _ = "STUB: not implemented"; return nil }

func (c *Context) HandlerName() string { _ = "STUB: not implemented"; return "" }

func (c *Context) HandlerNames() []string { _ = "STUB: not implemented"; return nil }

func (c *Context) Handler() HandlerFunc { _ = "STUB: not implemented"; return *new(HandlerFunc) }

func (c *Context) FullPath() string { _ = "STUB: not implemented"; return "" }

func (c *Context) Next() { _ = "STUB: not implemented"; return }

func (c *Context) IsAborted() bool { _ = "STUB: not implemented"; return false }

func (c *Context) Abort() { _ = "STUB: not implemented"; return }

func (c *Context) AbortWithStatus(code int) { _ = "STUB: not implemented"; return }

func (c *Context) AbortWithStatusPureJSON(code int, jsonObj any) { _ = "STUB: not implemented"; return }

func (c *Context) AbortWithStatusJSON(code int, jsonObj any) { _ = "STUB: not implemented"; return }

func (c *Context) AbortWithError(code int, err error) *Error { _ = "STUB: not implemented"; return nil }

func (c *Context) Error(err error) *Error { _ = "STUB: not implemented"; return nil }

func (c *Context) Set(key any, value any) { _ = "STUB: not implemented"; return }

func (c *Context) Get(key any) (value any, exists bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (c *Context) MustGet(key any) any { _ = "STUB: not implemented"; return *new(any) }

func getTyped[T any](c *Context, key any) (res T) { _ = "STUB: not implemented"; return *new(T) }

func (c *Context) GetString(key any) string { _ = "STUB: not implemented"; return "" }

func (c *Context) GetBool(key any) bool { _ = "STUB: not implemented"; return false }

func (c *Context) GetInt(key any) int { _ = "STUB: not implemented"; return 0 }

func (c *Context) GetInt8(key any) int8 { _ = "STUB: not implemented"; return 0 }

func (c *Context) GetInt16(key any) int16 { _ = "STUB: not implemented"; return 0 }

func (c *Context) GetInt32(key any) int32 { _ = "STUB: not implemented"; return 0 }

func (c *Context) GetInt64(key any) int64 { _ = "STUB: not implemented"; return 0 }

func (c *Context) GetUint(key any) uint { _ = "STUB: not implemented"; return 0 }

func (c *Context) GetUint8(key any) uint8 { _ = "STUB: not implemented"; return 0 }

func (c *Context) GetUint16(key any) uint16 { _ = "STUB: not implemented"; return 0 }

func (c *Context) GetUint32(key any) uint32 { _ = "STUB: not implemented"; return 0 }

func (c *Context) GetUint64(key any) uint64 { _ = "STUB: not implemented"; return 0 }

func (c *Context) GetFloat32(key any) float32 { _ = "STUB: not implemented"; return 0 }

func (c *Context) GetFloat64(key any) float64 { _ = "STUB: not implemented"; return 0 }

func (c *Context) GetTime(key any) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *Context) GetDuration(key any) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *Context) GetError(key any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) GetIntSlice(key any) []int { _ = "STUB: not implemented"; return nil }

func (c *Context) GetInt8Slice(key any) []int8 { _ = "STUB: not implemented"; return nil }

func (c *Context) GetInt16Slice(key any) []int16 { _ = "STUB: not implemented"; return nil }

func (c *Context) GetInt32Slice(key any) []int32 { _ = "STUB: not implemented"; return nil }

func (c *Context) GetInt64Slice(key any) []int64 { _ = "STUB: not implemented"; return nil }

func (c *Context) GetUintSlice(key any) []uint { _ = "STUB: not implemented"; return nil }

func (c *Context) GetUint8Slice(key any) []uint8 { _ = "STUB: not implemented"; return nil }

func (c *Context) GetUint16Slice(key any) []uint16 { _ = "STUB: not implemented"; return nil }

func (c *Context) GetUint32Slice(key any) []uint32 { _ = "STUB: not implemented"; return nil }

func (c *Context) GetUint64Slice(key any) []uint64 { _ = "STUB: not implemented"; return nil }

func (c *Context) GetFloat32Slice(key any) []float32 { _ = "STUB: not implemented"; return nil }

func (c *Context) GetFloat64Slice(key any) []float64 { _ = "STUB: not implemented"; return nil }

func (c *Context) GetStringSlice(key any) []string { _ = "STUB: not implemented"; return nil }

func (c *Context) GetErrorSlice(key any) []error { _ = "STUB: not implemented"; return nil }

func (c *Context) GetStringMap(key any) map[string]any { _ = "STUB: not implemented"; return nil }

func (c *Context) GetStringMapString(key any) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *Context) GetStringMapStringSlice(key any) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *Context) Delete(key any) { _ = "STUB: not implemented"; return }

func (c *Context) Param(key string) string { _ = "STUB: not implemented"; return "" }

func (c *Context) AddParam(key, value string) { _ = "STUB: not implemented"; return }

func (c *Context) Query(key string) (value string) { _ = "STUB: not implemented"; return "" }

func (c *Context) DefaultQuery(key, defaultValue string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Context) GetQuery(key string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (c *Context) QueryArray(key string) (values []string) { _ = "STUB: not implemented"; return nil }

func (c *Context) initQueryCache() { _ = "STUB: not implemented"; return }

func (c *Context) GetQueryArray(key string) (values []string, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *Context) QueryMap(key string) (dicts map[string]string) {
	_ = "STUB: not implemented"
	return nil
}

func (c *Context) GetQueryMap(key string) (map[string]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *Context) PostForm(key string) (value string) { _ = "STUB: not implemented"; return "" }

func (c *Context) DefaultPostForm(key, defaultValue string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Context) GetPostForm(key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (c *Context) PostFormArray(key string) (values []string) {
	_ = "STUB: not implemented"
	return nil
}

func (c *Context) initFormCache() { _ = "STUB: not implemented"; return }

func (c *Context) GetPostFormArray(key string) (values []string, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *Context) PostFormMap(key string) (dicts map[string]string) {
	_ = "STUB: not implemented"
	return nil
}

func (c *Context) GetPostFormMap(key string) (map[string]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func getMapFromFormData(m map[string][]string, key string) (map[string]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Context) MultipartForm() (*multipart.Form, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string, perm ...fs.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Context) Bind(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) BindJSON(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) BindXML(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) BindQuery(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) BindYAML(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) BindTOML(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) BindPlain(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) BindHeader(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) BindUri(obj any) error { _ = "STUB: not implemented"; return nil }

//nolint: errcheck

func (c *Context) MustBindWith(obj any, b binding.Binding) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint: errcheck

//nolint: errcheck

func (c *Context) ShouldBind(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) ShouldBindJSON(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) ShouldBindXML(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) ShouldBindQuery(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) ShouldBindYAML(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) ShouldBindTOML(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) ShouldBindPlain(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) ShouldBindHeader(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) ShouldBindUri(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) ShouldBindWith(obj any, b binding.Binding) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Context) ShouldBindBodyWith(obj any, bb binding.BindingBody) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *Context) ShouldBindBodyWithJSON(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) ShouldBindBodyWithXML(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) ShouldBindBodyWithYAML(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) ShouldBindBodyWithTOML(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) ShouldBindBodyWithPlain(obj any) error { _ = "STUB: not implemented"; return nil }

func (c *Context) ClientIP() string { _ = "STUB: not implemented"; return "" }

func (c *Context) RemoteIP() string { _ = "STUB: not implemented"; return "" }

func (c *Context) ContentType() string { _ = "STUB: not implemented"; return "" }

func (c *Context) IsWebsocket() bool { _ = "STUB: not implemented"; return false }

func (c *Context) Scheme() string { _ = "STUB: not implemented"; return "" }

func (c *Context) requestHeader(key string) string { _ = "STUB: not implemented"; return "" }

func bodyAllowedForStatus(status int) bool { _ = "STUB: not implemented"; return false }

func (c *Context) Status(code int) { _ = "STUB: not implemented"; return }

func (c *Context) Header(key, value string) { _ = "STUB: not implemented"; return }

func (c *Context) GetHeader(key string) string { _ = "STUB: not implemented"; return "" }

func (c *Context) GetRawData() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Context) SetSameSite(samesite http.SameSite) { _ = "STUB: not implemented"; return }

func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	_ = "STUB: not implemented"
	return
}

func (c *Context) SetCookieData(cookie *http.Cookie) { _ = "STUB: not implemented"; return }

func (c *Context) Cookie(name string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *Context) Render(code int, r render.Render) { _ = "STUB: not implemented"; return }

func (c *Context) HTML(code int, name string, obj any) { _ = "STUB: not implemented"; return }

func (c *Context) IndentedJSON(code int, obj any) { _ = "STUB: not implemented"; return }

func (c *Context) SecureJSON(code int, obj any) { _ = "STUB: not implemented"; return }

func (c *Context) JSONP(code int, obj any) { _ = "STUB: not implemented"; return }

func (c *Context) JSON(code int, obj any) { _ = "STUB: not implemented"; return }

func (c *Context) AsciiJSON(code int, obj any) { _ = "STUB: not implemented"; return }

func (c *Context) PureJSON(code int, obj any) { _ = "STUB: not implemented"; return }

func (c *Context) XML(code int, obj any) { _ = "STUB: not implemented"; return }

func (c *Context) PDF(code int, data []byte) { _ = "STUB: not implemented"; return }

func (c *Context) YAML(code int, obj any) { _ = "STUB: not implemented"; return }

func (c *Context) TOML(code int, obj any) { _ = "STUB: not implemented"; return }

func (c *Context) ProtoBuf(code int, obj any) { _ = "STUB: not implemented"; return }

func (c *Context) BSON(code int, obj any) { _ = "STUB: not implemented"; return }

func (c *Context) String(code int, format string, values ...any) { _ = "STUB: not implemented"; return }

func (c *Context) Redirect(code int, location string) { _ = "STUB: not implemented"; return }

func (c *Context) Data(code int, contentType string, data []byte) {
	_ = "STUB: not implemented"
	return
}

func (c *Context) DataFromReader(code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string) {
	_ = "STUB: not implemented"
	return
}

func (c *Context) File(filepath string) { _ = "STUB: not implemented"; return }

func (c *Context) FileFromFS(filepath string, fs http.FileSystem) {
	_ = "STUB: not implemented"
	return
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string { _ = "STUB: not implemented"; return "" }

func (c *Context) FileAttachment(filepath, filename string) { _ = "STUB: not implemented"; return }

func (c *Context) SSEvent(name string, message any) { _ = "STUB: not implemented"; return }

func (c *Context) Stream(step func(w io.Writer) bool) bool { _ = "STUB: not implemented"; return false }

type Negotiate struct {
	Offered      []string
	HTMLName     string
	HTMLData     any
	JSONData     any
	XMLData      any
	YAMLData     any
	Data         any
	TOMLData     any
	PROTOBUFData any
	BSONData     any
}

func (c *Context) Negotiate(code int, config Negotiate) { _ = "STUB: not implemented"; return }

//nolint: errcheck

func (c *Context) NegotiateFormat(offered ...string) string { _ = "STUB: not implemented"; return "" }

func (c *Context) SetAccepted(formats ...string) { _ = "STUB: not implemented"; return }

func (c *Context) hasRequestContext() bool { _ = "STUB: not implemented"; return false }

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func (c *Context) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (c *Context) Err() error { _ = "STUB: not implemented"; return nil }

func (c *Context) Value(key any) any { _ = "STUB: not implemented"; return *new(any) }
