// Package featureflag_header-modfication a plugin for traefik which adds request header.
package featureflag_header_modification

import (
	"context"
	"net/http"
	"os"
)

// Config the plugin configuration.
type Config struct {
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// XRequestStart a traefik plugin.
type FeatureflagHeaderModification struct {
	next http.Handler
	name string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &FeatureflagHeaderModification{
		next: next,
		name: name,
	}, nil
}

func (a *FeatureflagHeaderModification) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// TODO: implementation
	org := req.Header.Get("X-User-Org")
	os.Stdout.WriteString("traefik featuer flag header modification" + org)
	a.next.ServeHTTP(rw, req)
}
