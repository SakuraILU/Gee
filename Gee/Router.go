package gee

import (
	"fmt"
	log "gee/Log"
	trie "gee/Trie"
	"net/http"
	"strings"
)

type HandleFn func(*Context)
type Router struct {
	pattern_tire *trie.Trie
	handlers     map[string]HandleFn
}

func newRouter() (r *Router) {
	return &Router{
		pattern_tire: trie.NewTrie(),
		handlers:     make(map[string]HandleFn),
	}
}

func (r *Router) addHandler(method, pattern string, fn HandleFn) {
	r.pattern_tire.Insert(pattern)
	key := keyGen(method, pattern)
	log.Infof("add router key %v", key)
	r.handlers[key] = fn
}

func (r *Router) handle(c *Context) (err error) {
	pattern := r.pattern_tire.Search(c.Path)
	if pattern == "" {
		err = fmt.Errorf("the pattern matched url %s is not found", c.Path)
		c.String(http.StatusNotFound, err.Error())
		log.Error(err)
		return
	}
	key := keyGen(c.Method, pattern)
	handler, ok := r.handlers[key]
	if !ok {
		err = fmt.Errorf("the handler for pattern %s is not found", pattern)
		c.String(http.StatusNotFound, err.Error())
		log.Error(err)
	}

	c.Params = parseParams(pattern, c.Path)

	handler(c)
	return
}

func keyGen(method, pattern string) (key string) {
	return method + "-" + pattern
}

func parseParams(pattern, url string) (params map[string]string) {
	params = make(map[string]string)

	patparts := strings.Split(pattern, "/")
	urlparts := strings.Split(url, "/")
	for i, patpart := range patparts {
		if patpart == "" {
			continue
		}

		switch patpart[0] {
		case ':':
			params[patpart[1:]] = urlparts[i]
		case '*':
			params[patpart[1:]] = strings.Join(urlparts[i:], "/")
		}
	}
	return
}
