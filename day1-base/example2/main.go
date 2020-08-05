package main

import (
	"ebb"
	"fmt"
	"net/http"
)

func main(){
	r := ebb.New()

	r.GET("/",func(w http.ResponseWriter,req *http.Request){
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	
	r.Run(":8080")
}