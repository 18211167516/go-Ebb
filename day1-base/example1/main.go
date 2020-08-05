package main

import (
	"fmt"
	"net/http"
	"log"
)

func indexHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func helloHandler(w http.ResponseWriter,r *http.Request){
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

func main(){
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/hello",helloHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}