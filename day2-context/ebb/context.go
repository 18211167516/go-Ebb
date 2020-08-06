package ebb


import (
	"net/http"
	"fmt"
	"encoding/json"
)

type H map[string]interface{}

type Context struct{
	//write and request
	Writer http.ResponseWriter
	Request *http.Request
	//request info
	Method string
	Path string
	//response
	HttpCode int
}


func newContext(w http.ResponseWriter,r *http.Request) *Context{
	context := &Context{
		Writer:w,
		Request:r,
		Path:   r.URL.Path,
		Method: r.Method,
	}

	return context
} 

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.HttpCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) Write(data []byte){
	c.Writer.Write(data)
}


func (c *Context) String(code int,message string,v ...interface{}){
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Write([]byte(fmt.Sprintf(message, v...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	data,err:= json.Marshal(obj)
	if err!=nil {
		http.Error(c.Writer, err.Error(), 500)
	}
	c.Write(data)
}

func (c *Context) HTML(code int,html string){
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Write([]byte(html))
}