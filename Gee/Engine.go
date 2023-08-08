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
	key := "GET" + "-" + pattern
	e.router.addRouter(key, fn)
}

func (e *Engine) POST(pattern string, fn HandleFn) {
	key := "POST" + "-" + pattern
	e.router.addRouter(key, fn)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)

	key := c.Method + "-" + c.Path
	handler, err := e.router.getRouter(key)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	handler(c)
}

func (e *Engine) Run(addr string) (err error) {
	log.Infof("Service is listening at %v", addr)
	return http.ListenAndServe(addr, e)
}
