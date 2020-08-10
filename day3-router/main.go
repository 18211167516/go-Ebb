package main

import (
	"ebb"
)


func main(){
	r := ebb.New()

	r.GET("/hello/:lang/doc", func(c *ebb.Context) {
		c.HTML(200, "<h1>Hello ebb</h1>")
	})
	r.GET("/hello/:name", func(c *ebb.Context) {
		c.String(200, "hello %s, you from %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:id/:name", func(c *ebb.Context) {
		
		c.String(200, "hello %s, you from %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login/*filepath", func(c *ebb.Context) {
		c.JSON(200, ebb.H{
			"name": c.PostForm("name"),
		})
	})

	
	r.Run(":8080")
}