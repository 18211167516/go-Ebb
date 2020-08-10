
package ebb

import (
	"net/http"
	"testing"
)

func BenchmarkOneRoute(B *testing.B) {
	router := New()
	router.GET("/ping", func(c *Context) {})
	runRequest(B, router, "GET", "/ping")
}

func BenchmarkOneRouteJSON(B *testing.B) {
	router := New()
	data := struct {
		Status string `json:"status"`
	}{"ok"}
	router.GET("/json", func(c *Context) {
		c.JSON(http.StatusOK, data)
	})
	runRequest(B, router, "GET", "/json")
}

func BenchmarkOneRouteHTML(B *testing.B) {
	router := New()

	router.GET("/html", func(c *Context) {
		c.HTML(http.StatusOK, "<h1>Hello ebb</h1>")
	})
	runRequest(B, router, "GET", "/html")
}

func BenchmarkOneRouteString(B *testing.B) {
	router := New()
	router.GET("/text", func(c *Context) {
		c.String(http.StatusOK, "this is a plain text")
	})
	runRequest(B, router, "GET", "/text")
}


func Benchmark404(B *testing.B) {
	router := New()
	router.GET("/something", func(c *Context) {})
	runRequest(B, router, "GET", "/ping")
}


