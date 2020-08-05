package ebb

import (
	"net/http"
	"log"
)

type HandlerFunc func(w http.ResponseWriter,req *http.Request)


type Engine struct{
	router map[string]HandlerFunc
}

func New() *Engine{
	engine := &Engine{
		router : make(map[string]HandlerFunc),
	}
	return engine
}

func (engine *Engine) addRoute(mothod string,pattern string,handler HandlerFunc){
	key := mothod+"-"+pattern
	engine.router[key] = handler
}

func (engine *Engine) GET(pattern string,handler HandlerFunc){
	engine.addRoute("GET",pattern,handler)
}

func (engine *Engine) POST(pattern string,handler HandlerFunc){
	engine.addRoute("POST",pattern,handler)
}

func (engine *Engine) Run(addr string) (err error){
	return http.ListenAndServe(addr,engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter,req *http.Request){
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		log.Printf("404 NOT FOUND: %s\n", req.URL)
	}
}
