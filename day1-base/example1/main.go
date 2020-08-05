package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func helloHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"URL.PATH=%q\n Header=%q\n",r.URL.Path,r.Header)
}

func main(){
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/hello",helloHandler)
	http.ListenAndServe(":8080",nil)
}