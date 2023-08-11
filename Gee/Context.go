package gee

import (
	"encoding/json"
	"fmt"
	log "gee/Log"
	"net/http"
)

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Method     string
	Path       string
	StatusCode int
	// url params
	Params map[string]string
	// middleware and handler
	handlers []HandleFn
	index    int
}

func newContext(w http.ResponseWriter, req *http.Request) (c *Context) {
	return &Context{
		Writer:     w,
		Req:        req,
		Method:     req.Method,
		Path:       req.URL.Path,
		StatusCode: http.StatusOK,
		Params:     make(map[string]string),
		handlers:   make([]HandleFn, 0),
		index:      -1,
	}
}

func (c *Context) String(code int, format string, vs ...any) {
	c.SetCode(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, vs...)))
}

func (c *Context) HTML(code int, html string) {
	c.SetCode(code)
	c.Writer.Header().Set("Content-Type", "application/html")
	c.Writer.Write([]byte(html + "\n"))
}

func (c *Context) JSON(code int, obj any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.SetCode(code)
	if err := json.NewEncoder(c.Writer).Encode(obj); err != nil {
		log.Error(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Context) DATA(code int, data []byte) {
	c.SetCode(code)
	c.Writer.Write(data)
}

func (c *Context) Fail(code int, err string) {
	// if err occurs, stop subsequential calls of handlers (middlewares or handler)
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

func (c *Context) SetCode(code int) {
	c.Writer.WriteHeader(code)
	c.StatusCode = code
}

// get the value of the key in the GET url
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// get the value corresponding to the key (key, value) pair in the POST request
func (c *Context) PostForm(key string) string {
	return c.Req.PostFormValue(key)
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
	}
}
