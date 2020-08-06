package ebb

type router struct{
	handlers map[string]HandlerFunc
}

func newRouter() *router{
	return &router{handlers:make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string,pattern string,handler HandlerFunc){
	key := method+"-"+pattern
	r.handlers[key] = handler
}

