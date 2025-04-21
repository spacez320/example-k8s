package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpResponseOk(t *testing.T) {
	tests := map[string]struct {
		method string
		path   string
		status int
	}{
		"GET /": {
			method: "GET",
			path:   "",
			status: http.StatusOK,
		},
		"POST /foo/bar": {
			method: "POST",
			path:   "foo/bar",
			status: http.StatusOK,
		},
	}

	for name, test := range tests {
		req := httptest.NewRequest(test.method, fmt.Sprintf("/%s", test.path), nil)
		rec := httptest.NewRecorder()

		httpResponseOk(rec, req)
		resp := rec.Result()
		defer resp.Body.Close()

		if resp.StatusCode != test.status {
			t.Errorf("Error: %s HTTP status not OK", name)
		}
	}
}
