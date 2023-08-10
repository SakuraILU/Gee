package gee

import (
	log "gee/Log"
	"time"
)

func Logger() HandleFn {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Infof("[%d] %s within %v", c.StatusCode, c.Path, time.Since(t))
	}
}
