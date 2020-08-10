package main

import (
	"ebb"
)


func main(){
	r := ebb.New()

	r.GET("/index", func(c *ebb.Context) {
		c.HTML(200, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *ebb.Context) {
			c.HTML(200, "<h1>Hello ebb</h1>")
		})

		v1.GET("/hello", func(c *ebb.Context) {
			c.String(200, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *ebb.Context) {
			// expect /hello/ebbktutu
			c.String(200, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *ebb.Context) {
			c.JSON(200, ebb.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}	
	r.Run(":8080")
}