package main

import (
	"ebb"
)


func main(){
	r := ebb.New()

	r.GET("/", func(c *ebb.Context) {
		c.HTML(200, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *ebb.Context) {
		// expect /hello?name=geektutu
		c.String(200, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *ebb.Context) {
		c.JSON(200, ebb.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	
	r.Run(":8080")
}