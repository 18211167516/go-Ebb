package ebb

import (
	"net/http"
)

//定义框架请求处理方法
type HandlerFunc func(*Context)

//核心结构体
type Engine struct{
	*RouterGroup //v4新增  顶级路由组
	router *router 
	groups []*RouterGroup // v4新增
}

//实例化结构体
func New() *Engine{
	engine := &Engine{
		router : newRouter(),
	}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

//添加到结构体路由
func (engine *Engine) addRoute(mothod string,pattern string,handler HandlerFunc){
	engine.router.addRoute(mothod,pattern,handler)
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
	c := newContext(w, req)
	engine.handleHTTPRequest(c)
}

//v2 新增
func (engine *Engine) handleHTTPRequest(c *Context){
	node,params := engine.router.getRoute(c.Method,c.Path)
	if node != nil {
		c.Params = params
		key := c.Method + "-" + node.pattern
		//调用实际的方法
		engine.router.handlers[key](c)
	}else{
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
