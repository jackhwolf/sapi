package sapi

import (
	"context"
)

// Route maps a handler to a url and method(s)
type Route struct {
	Method  string
	Path    string
	Handler func(context.Context, Payload) *HandlerReturn
}

// NewRoute is the entry point for making a new route
func NewRoute(handler func(context.Context, Payload) *HandlerReturn, path string, methods string) *Route {
	r := &Route{
		Method:  methods,
		Path:    path,
		Handler: handler,
	}
	return r
}
