package main

import (
	"fmt"
	gee "gee/Gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.Use(gee.Logger()) // global midlleware
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/*name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
			err := fmt.Errorf("a error occurs in /v2/hello/%s", c.Param("name"))
			panic(err)
		})
	}

	v2.Static("/asserts", "Filesystem")
	r.Run(":9999")
}
