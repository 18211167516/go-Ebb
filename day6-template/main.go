package main

import (
	"ebb"
	"time"
	"fmt"
)

type student struct {
	Name string
	Age  int8
}

func FormatDate(t time.Time) string {
	year, month, day := t.Date()
	Hour,Minute,Second := t.Hour(),t.Minute(),t.Second()
	return fmt.Sprintf("%d-%02d-%02d %d:%d:%d", year, month, day,Hour,Minute,Second)
}

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
	r.Static("/assets", "./static")
	r.SetFuncMap("formatDate",FormatDate)
	r.LoadHTMLGlob("templates/*")
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}

	r.GET("/students", func(c *ebb.Context) {
		c.HTML(200, "students.tmpl", ebb.H{
			"title":  "ebb  students",
			"now":   time.Now(),
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.Run(":8080")
}