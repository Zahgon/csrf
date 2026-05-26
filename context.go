//go:build go1.7
// +build go1.7

package csrf

import (
	"net/http"
)

func contextGet(r *http.Request, key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func contextSave(r *http.Request, key string, val interface{}) *http.Request {
	_ = "STUB: not implemented"
	return nil
}

// nolint:staticcheck

func contextClear(r *http.Request) {
	_ = "STUB: not implemented"
	// no-op for go1.7+
	return
}
