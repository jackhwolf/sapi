package sapi

import (
	"context"
	"net/http"
	"testing"
)

// TestNewRouter
func TestNewRouter(t *testing.T) {
	// create new router with url prefix `/prefix`
	NewRouter("/prefix/")
}

// TestAddPath
func TestAddPath(t *testing.T) {
	// create new router with url prefix `/prefix`
	rtr := NewRouter("/prefix/")

	// assign sampleData function to GET /prefix/sample
	rtr.AddRoute(func(ctx context.Context, payload Payload) *HandlerReturn {
		sample := struct {
			Message string
			Time    int
		}{"Hello", 123}
		return &HandlerReturn{sample, 200, nil}
	}, "/sample", http.MethodGet)
}
