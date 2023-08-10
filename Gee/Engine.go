package gee

import (
	log "gee/Log"
	"net/http"
)

type Engine struct {
	router *Router
}

func New() (e *Engine) {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) GET(pattern string, fn HandleFn) {
	e.router.addHandler("GET", pattern, fn)
}

func (e *Engine) POST(pattern string, fn HandleFn) {
	e.router.addHandler("POST", pattern, fn)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)

	if err := e.router.handle(c); err != nil {
		return
	}
}

func (e *Engine) Run(addr string) (err error) {
	log.Infof("Service is listening at %v", addr)
	return http.ListenAndServe(addr, e)
}
