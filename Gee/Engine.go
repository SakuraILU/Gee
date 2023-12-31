package gee

import (
	log "gee/Log"
	"net/http"
	"strings"
)

type Engine struct {
	router *Router

	groups      []*Group
	middlewares []HandleFn
}

func New() (e *Engine) {
	return &Engine{
		router:      newRouter(),
		groups:      make([]*Group, 0),
		middlewares: make([]HandleFn, 0),
	}
}

func Default() (e *Engine) {
	e = &Engine{
		router:      newRouter(),
		groups:      make([]*Group, 0),
		middlewares: make([]HandleFn, 0),
	}
	e.Use(Logger(), Recovery())

	return
}

func (e *Engine) GET(pattern string, fn HandleFn) {
	e.router.addHandler("GET", pattern, fn)
}

func (e *Engine) POST(pattern string, fn HandleFn) {
	e.router.addHandler("POST", pattern, fn)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)

	c.handlers = append(c.handlers, e.middlewares...)
	for _, g := range e.groups {
		if strings.HasPrefix(c.Path, g.prefix) {
			c.handlers = append(c.handlers, g.middlewares...)
		}
	}

	if err := e.router.handle(c); err != nil {
		return
	}
}

func (e *Engine) Run(addr string) (err error) {
	log.Infof("Service is listening at %v", addr)
	return http.ListenAndServe(addr, e)
}

func (e *Engine) Group(prefix string) (g *Group) {
	return newGroup(""+prefix, e)
}

func (e *Engine) Use(middleware ...HandleFn) {
	e.middlewares = append(e.middlewares, middleware...)
}

func (e *Engine) Static(url_dir_path, dir_path string) {
	pattern := url_dir_path + "/*filerelpath"
	dir := http.Dir(dir_path)
	fs := http.FileServer(dir)
	fs = http.StripPrefix(url_dir_path, fs)
	e.GET(pattern, func(ctx *Context) {
		fs.ServeHTTP(ctx.Writer, ctx.Req)
	})
}
