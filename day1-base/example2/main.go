package main

import (
	"ebb"
	"fmt"
	"net/http"
)

func main(){
	r := ebb.New()

	r.GET("/index",func(w http.ResponseWriter,req *http.Request){
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, req.Header)
	})
	
	r.Run(":8080")
}