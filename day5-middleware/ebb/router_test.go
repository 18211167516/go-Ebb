package ebb


import (
	"testing"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestGetRoute(t *testing.T) {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	n, params := r.getRoute("GET", "/index/baibai")

	if assert.NotNil(t, n,"404 not found ") {
		assert.Equal(t,n.pattern,"/hello/:name","should match /hello/:name")
		assert.Equal(t,params["name"],"baibai","name should be equel to 'baibai'")
	}

}
func TestParsePattern(t *testing.T) {
	assert.Equal(t,parsePattern("/p/:name"),[]string{"p",":name"},"not parsePattern :name")
	assert.Equal(t,parsePattern("/p/*"),[]string{"p","*"},"not parsePattern *")
	assert.Equal(t,parsePattern("/p/*name/*"),[]string{"p","*name"},"parsePattern not truncation")
}
func TestRouter(t *testing.T){
	r := New()

	r.POST("/login/*filepath", func(c *Context) {
		c.JSON(200, H{
			"name": c.PostForm("name"),
		})
	})

	param := `{"name":"56789","state":3}`

	w := PerformRequest("POST","/login/123213?name=1233",bytes.NewBufferString(param),r)

	s := struct{
		Name string `json:"name"`
	}{}
	json.Unmarshal([]byte(w.Body.String()),&s)

	fmt.Printf("%+v",s)
	assert.Equal(t,s.Name,"1233","PostForm error")
}



func TestMiddleware(t *testing.T){
	r := New()

	r.Use(Logger())
	r.POST("/login/*filepath",Logger(),func(c *Context) {
		c.JSON(200, H{
			"name": c.PostForm("name"),
		})
	})

	param := `{"name":"56789","state":3}`

	w := PerformRequest("POST","/login/123213?name=1233",bytes.NewBufferString(param),r)

	s := struct{
		Name string `json:"name"`
	}{}
	json.Unmarshal([]byte(w.Body.String()),&s)

	assert.Equal(t,s.Name,"1233","PostForm error")
}

func TestRecovery(t *testing.T){
	r := New()

	v1 := r.Group("/v1")
	v1.Use(Logger(),Recovery())
	{
		v1.GET("/panic",func(c *Context) {
			panic("err")
		})
	}

	r.POST("/login/*name",func(c *Context) {
		c.JSON(200, H{
			"name": c.Param("name"),
		})
	})

	param := `{"name":"baibai"}`;
	_ = PerformRequest("GET","/v1/panic",bytes.NewBufferString(param),r)

	_ = PerformRequest("POST","/login/123",bytes.NewBufferString(param),r)

}