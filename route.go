package sapi

import (
	"context"
)

// Route maps a handler to a url and method(s)
type Route struct {
	Methods []string
	Path    string
	Handler func(context.Context, Payload) (interface{}, int, error)
}

// NewRoute is the entry point for making a new route
func NewRoute(handler func(context.Context, Payload) (interface{}, int, error), path string, methods ...string) *Route {
	r := &Route{
		Methods: methods,
		Path:    path,
		Handler: handler,
	}
	return r
}

// RouteHasMethod checks if the route found by URL search has the requested method
func (rt *Route) RouteHasMethod(requestedMethod string) bool {
	for _, s := range rt.Methods {
		if s == requestedMethod {
			return true
		}
	}
	return false
}
