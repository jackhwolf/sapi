package sapi

import (
	"context"
	"net/http"
	"testing"
)

func TestNewRouter(t *testing.T) {
	// create new router with url prefix `/prefix`
	NewRouter("/prefix/")
}

func TestAddPath(t *testing.T) {
	// create new router with url prefix `/prefix`
	rtr := NewRouter("/prefix/")

	// assign sampleData function to GET /prefix/sample
	rtr.AddRoute(func(ctx context.Context, payload Payload) (interface{}, int, error) {
		sample := struct {
			Message string
			Time    int
		}{"Hello", 123}
		return sample, 200, nil
	}, "/sample", http.MethodGet)

}
