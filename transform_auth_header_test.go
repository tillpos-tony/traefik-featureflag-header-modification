package transform_auth_header_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hitz-group/traefik-copy-auth-header"
)

func TestXRequestStart(t *testing.T) {
	cfg := transform_auth_header.CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := transform_auth_header.New(ctx, next, cfg, "transform_auth_header")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)
}
