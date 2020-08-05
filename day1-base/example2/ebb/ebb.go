package ebb

import (
	"net/http"
	"log"
)

//定义框架请求处理方法
type HandlerFunc func(w http.ResponseWriter,req *http.Request)


//核心结构体
type Engine struct{
	router map[string]HandlerFunc //简单使用map记录路由信息
}

//实例化结构体
func New() *Engine{
	engine := &Engine{
		router : make(map[string]HandlerFunc),
	}
	return engine
}

//添加到结构体路由
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

//启动服务
func (engine *Engine) Run(addr string) (err error){
	return http.ListenAndServe(addr,engine)
}

//engine 实现ServeHTTP接口（所有的请求都会走到这）
//查找是否路由映射表存在，如果存在则调用，否则返回404
func (engine *Engine) ServeHTTP(w http.ResponseWriter,req *http.Request){
	
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		log.Printf("404 NOT FOUND: %s\n", req.URL)
	}
}
