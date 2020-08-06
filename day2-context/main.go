package main

import (
	"ebb"
)


func main(){
	r := ebb.New()

	r.GET("/", func(c *ebb.Context) {
		c.HTML(200, "<h1>Hello ebb</h1>")
	})
	r.GET("/hello", func(c *ebb.Context) {
		c.String(200, "hello %s, you from %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *ebb.Context) {
		c.JSON(200, ebb.H{
			"name": c.PostForm("name"),
		})
	})

	
	r.Run(":8080")
}