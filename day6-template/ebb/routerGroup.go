package ebb

import (
	"path"
	"net/http"
)
type HandlersChain []HandlerFunc

type RouterGroup struct {
	prefix      string
	middlewares HandlersChain // support middleware
	parent      *RouterGroup  // support nesting
	engine      *Engine       // all groups share a Engine instance
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,//支持分组嵌套
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

//添加中间件 v5新增
func (group *RouterGroup) Use(middlewares ...HandlerFunc){
	group.middlewares = append(group.middlewares, middlewares...)
}

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	finalSize := len(group.middlewares) + len(handlers)
	mergedHandlers := make(HandlersChain, finalSize)
	copy(mergedHandlers, group.middlewares)
	copy(mergedHandlers[len(group.middlewares):], handlers)
	return mergedHandlers
}

func (group *RouterGroup) addRoute(method string, comp string, handler ...HandlerFunc) {
	pattern := group.prefix + comp
	handlers := group.combineHandlers(handler)
	group.engine.router.addRoute(method, pattern, handlers)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler ...HandlerFunc) {
	group.addRoute("GET", pattern, handler...)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler ...HandlerFunc) {
	group.addRoute("POST", pattern, handler...)
}

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := joinPaths(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		f, err := fs.Open(file);
		defer f.Close()
		if  err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}

// serve static files
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := joinPaths(relativePath, "/*filepath")
	// Register GET handlers
	group.GET(urlPattern, handler)
}


func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}
	finalPath := path.Join(absolutePath, relativePath)
	return finalPath
}