package trie

type node struct {
	pattern  string
	name     string
	iswild   bool
	children []*node
}

func (n *node) childrenMatched(name string) (children []*node) {
	for _, child := range n.children {
		if child.name == name || child.iswild {
			children = append(children, child)
		}
	}
	return
}

func (n *node) childMatchedPrecious(name string) (child *node) {
	for _, child := range n.children {
		if child.name == name {
			return child
		}
	}
	return
}

// must gurantee the correctness of parts:
//
//	only parts[-1] can start with * representing *filepath
func (n *node) insert(parts []string, pattern string) {
	if len(parts) == 0 {
		n.pattern = pattern
		return
	}

	name := parts[0]
	child := n.childMatchedPrecious(name)
	if child == nil {
		child = &node{
			name:     name,
			iswild:   (name[0] == '*' || name[0] == ':'),
			children: make([]*node, 0),
		}
		n.children = append(n.children, child)
	}
	child.insert(parts[1:], pattern)
}

func (n *node) search(parts []string) *node {
	// log.Infof("search %s", n.name)
	if len(parts) == 0 {
		if n.pattern != "" {
			return n
		} else {
			return nil
		}
	}

	name := parts[0]
	children := n.childrenMatched(name)
	if len(children) == 0 {
		return nil
	}

	for _, child := range children {
		if child.name == name || (child.iswild && child.name[0] == ':') {
			if res := child.search(parts[1:]); res != nil {
				return res
			}
		} else if child.iswild && child.name[0] == '*' {
			return child
		}
	}
	return nil
}
