package gin

import (
	"net/http"
)

func CreateTestContext(w http.ResponseWriter) (c *Context, r *Engine) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateTestContextOnly(w http.ResponseWriter, r *Engine) (c *Context) {
	_ = "STUB: not implemented"
	return nil
}

func waitForServerReady(url string, maxAttempts int) error { _ = "STUB: not implemented"; return nil }
