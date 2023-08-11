package gee

import (
	"fmt"
	log "gee/Log"
	"net/http"
	"runtime"
	"strings"
	"time"
)

const NTRACE int = 32
const NIGNORE int = 3

func Logger() HandleFn {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Infof("[%d] %s within %v", c.StatusCode, c.Path, time.Since(t))
	}
}

func Recovery() HandleFn {
	return func(ctx *Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Error(traceMsg())
				ctx.Fail(http.StatusInternalServerError, fmt.Sprint(err))
			}
		}()
		ctx.Next()
	}
}

func traceMsg() string {
	msg := &strings.Builder{}
	msg.WriteString("Trace Back: ")
	pcs := make([]uintptr, NTRACE)
	// ignore three callers:
	//		panic()  -->  defer() (in recovery middleware) -->  traceMsg()
	runtime.Callers(NIGNORE, pcs)
	frames := runtime.CallersFrames(pcs)
	for {
		if frame, ok := frames.Next(); ok {
			msg.WriteString(fmt.Sprintf("[%s] %s:%d\n", frame.Function, frame.File, frame.Line))
		} else {
			break
		}
	}

	return msg.String()
}
