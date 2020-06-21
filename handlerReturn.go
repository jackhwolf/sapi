package sapi

// HandlerReturn wraps everything to return
type HandlerReturn struct {
	Body       interface{}
	StatusCode int
	Err        error
}
