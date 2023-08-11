package gee

type Group struct {
	prefix      string
	engine      *Engine
	middlewares []HandleFn
}

func newGroup(prefix string, engine *Engine) (g *Group) {
	g = &Group{
		prefix:      prefix,
		engine:      engine,
		middlewares: make([]HandleFn, 0),
	}
	engine.groups = append(engine.groups, g)
	return
}

func (g *Group) Group(prefix string) (ng *Group) {
	ng = newGroup(g.prefix+prefix, g.engine)
	return
}

func (g *Group) GET(pattern string, handler HandleFn) {
	g.engine.GET(g.prefix+pattern, handler)
}

func (g *Group) POST(pattern string, handler HandleFn) {
	g.engine.POST(g.prefix+pattern, handler)
}

func (g *Group) Use(middleware ...HandleFn) {
	g.middlewares = append(g.middlewares, middleware...)
}

func (g *Group) Static(url_dir_path, dir_path string) {
	g.engine.Static(g.prefix+url_dir_path, dir_path)
}
