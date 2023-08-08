package gee

import (
	log "gee/Log"
)

type HandleFn func(*Context)
type Router struct {
	handlers map[string]HandleFn
}

func newRouter() (r *Router) {
	return &Router{
		handlers: make(map[string]HandleFn),
	}
}

func (r *Router) addRouter(pattern string, fn HandleFn) (err error) {
	log.Infof("add router key %v", pattern)
	r.handlers[pattern] = fn
	return nil
}

func (r *Router) getRouter(pattern string) (handler HandleFn, err error) {
	handler, ok := r.handlers[pattern]
	if !ok {
		log.Error("handler pattern %v is not found", pattern)
	}
	return
}
