package sapi

import (
	"context"
	"encoding/json"
	"errors"
	"log"
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
	trimPref := strings.TrimRight(pref, "/")
	rm := make(map[string]*Route, 0)
	r := &Router{
		Prefix:   trimPref,
		RouteMap: rm,
	}
	return r
}

func (rtr *Router) makeMethodPathKey(method, path string) string {
	return method + ":" + path
}

// LookupRoute will lookup the requested route
func (rtr *Router) LookupRoute(method, path string) (*Route, error) {
	key := rtr.makeMethodPathKey(method, path)
	routeMapEntry, routeMapOk := rtr.RouteMap[key]
	if !routeMapOk {
		return &Route{}, errors.New(method + " " + path + " not defined")
	}
	return routeMapEntry, nil
}

// addRoute will try to add a route
func (rtr *Router) addRoute(handler func(context.Context, Payload) *HandlerReturn, path string, method string) error {
	path = rtr.Prefix + path
	_, err := rtr.LookupRoute(method, path)
	if err != nil {
		return errors.New(method + " " + path + " already defined")
	}
	key := rtr.makeMethodPathKey(method, path)
	rtr.RouteMap[key] = NewRoute(handler, path, method)
	return nil
}

// AddRoute will try to add a route
func (rtr *Router) AddRoute(handler func(context.Context, Payload) *HandlerReturn, path string, methods ...string) error {
	path = rtr.Prefix + path
	for _, m := range methods {
		err := rtr.addRoute(handler, path, m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (rtr *Router) logHit(payload Payload) {
	log.Println(payload.HTTPMethod, payload.Path)
}

// HandleLambda should be passed to lambda.start as the entry point for requests
func (rtr *Router) HandleLambda(ctx context.Context, payload Payload) (Response, error) {

	rtr.logHit(payload)

	response := DefaultResponse()

	route, err := rtr.LookupRoute(payload.HTTPMethod, payload.Path)
	if err != nil {
		response.Body = err.Error()
		return *response, nil
	}

	handlerReturn := route.Handler(ctx, payload)
	var respBodyStr string

	if handlerReturn.Err != nil {
		response.Body = handlerReturn.Err.Error()
	} else {
		respBodyByte, err := json.Marshal(handlerReturn.Body)
		if err != nil {
			handlerReturn.StatusCode = http.StatusInternalServerError
			respBodyStr = http.StatusText(handlerReturn.StatusCode)
		} else {
			respBodyStr = string(respBodyByte)
			response.Headers.ContentType = "application/json"
		}
	}

	response.StatusCode = handlerReturn.StatusCode
	response.StatusDescription = http.StatusText(response.StatusCode)
	response.Body = respBodyStr
	return *response, nil
}
