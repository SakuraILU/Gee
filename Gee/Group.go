package gee

type Group struct {
	prefix string
	engine *Engine
}

func (g *Group) Group(prefix string) *Group {
	return &Group{
		prefix: g.prefix + prefix,
		engine: g.engine,
	}
}

func (g *Group) GET(pattern string, handler HandleFn) {
	g.engine.GET(g.prefix+pattern, handler)
}

func (g *Group) POST(pattern string, handler HandleFn) {
	g.engine.POST(g.prefix+pattern, handler)
}
