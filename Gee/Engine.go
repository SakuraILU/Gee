package gee

import (
	"fmt"
	log "gee/Log"
	"net/http"
)

type HandleFn func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandleFn
}

func New() (e *Engine) {
	return &Engine{
		router: make(map[string]HandleFn),
	}
}

func (e *Engine) addRouter(key string, fn HandleFn) {
	log.Infof("add router key %v", key)
	e.router[key] = fn
}

func (e *Engine) GET(pattern string, fn HandleFn) {
	key := "GET" + "-" + pattern
	e.addRouter(key, fn)
}

func (e *Engine) POST(pattern string, fn HandleFn) {
	key := "POST" + "-" + pattern
	e.addRouter(key, fn)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	handler, ok := e.router[key]
	if !ok {
		err := fmt.Errorf("key %v is not found", key)
		log.Errorf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	handler(w, req)
}

func (e *Engine) Run(addr string) (err error) {
	log.Infof("Service is listening at %v", addr)
	return http.ListenAndServe(addr, e)
}
