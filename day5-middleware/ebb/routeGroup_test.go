package ebb


import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRouteGroup(t *testing.T) {
	r := New()
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *Context) {
			// expect /hello/ebbktutu
			c.String(200, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *Context) {
			c.JSON(200, H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}
	n, params := r.router.getRoute("GET", "v2/hello/baibai")
	if assert.NotNil(t, n,"404 not found ") {
		assert.Equal(t,n.pattern,"/v2/hello/:name","should match /hello/:name")
		assert.Equal(t,params["name"],"baibai","name should be equel to 'baibai'")
	}

}