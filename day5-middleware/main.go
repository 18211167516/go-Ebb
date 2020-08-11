package main

import (
	"ebb"
)


func main(){
	r := ebb.Default()
	r.GET("/panic",func(c *ebb.Context) {
		panic("err")
	})

	r.POST("/login/*name",func(c *ebb.Context) {
		c.JSON(200, ebb.H{
			"name": c.Param("name"),
		})
	})
	r.Run(":8080")
}