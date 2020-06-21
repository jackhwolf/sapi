package sapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// Router will map requests to handlers
type Router struct {
	Prefix   string
	RouteMap map[string]*Route
}

// NewRouter is the entry point for making a new router
func NewRouter(pref string) *Router {
	pref = strings.TrimRight(pref, "/")
	rm := make(map[string]*Route, 0)
	r := &Router{
		Prefix:   pref,
		RouteMap: rm,
	}
	return r
}

// LookupRoute will lookup the requested route
func (rtr *Router) LookupRoute(method, path string) (*Route, error) {
	routeMapEntry, routeMapOk := rtr.RouteMap[path]
	if !routeMapOk {
		return &Route{}, errors.New(path + " not defined")
	}
	if !routeMapEntry.RouteHasMethod(method) {
		return &Route{}, errors.New(method + " not defined for path " + path)
	}
	return routeMapEntry, nil
}

// AddRoute will try to add a route
func (rtr *Router) AddRoute(handler func(context.Context, Payload) (interface{}, int, error), path string, methods ...string) error {
	path = rtr.Prefix + path
	for _, m := range methods {
		_, err := rtr.LookupRoute(m, path)
		if err == nil {
			return errors.New(m + ":" + path + " already defined")
		}
	}
	rtr.RouteMap[path] = NewRoute(handler, path, methods...)
	return nil
}

// HandleLambda should be passed to lambda.start as the entry point for requests
func (rtr *Router) HandleLambda(ctx context.Context, payload Payload) (Response, error) {
	response := &Response{}
	response.Headers.ContentType = "text/html"
	response.StatusCode = http.StatusBadRequest
	response.StatusDescription = http.StatusText(http.StatusBadRequest)
	route, err := rtr.LookupRoute(payload.HTTPMethod, payload.Path)
	if err != nil {
		response.Body = err.Error()
		return *response, nil
	}
	respBody, respStatusInt, respErr := route.Handler(ctx, payload)
	var respBodyStr string
	if respErr != nil {
		response.Body = respErr.Error()
	} else {
		respBodyByte, err := json.Marshal(respBody)
		respBodyStr = string(respBodyByte)
		if err != nil {
			respStatusInt = http.StatusInternalServerError
			respBodyStr = http.StatusText(respStatusInt)
		}
	}
	response.StatusCode = respStatusInt
	response.StatusDescription = http.StatusText(respStatusInt)
	response.Body = respBodyStr
	return *response, nil
}
