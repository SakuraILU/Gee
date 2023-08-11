package trie

import (
	"strings"
)

type Trie struct {
	root *node
}

func NewTrie() (t *Trie) {
	return &Trie{
		root: &node{
			pattern:  "/",
			children: make([]*node, 0),
		},
	}
}

func (t *Trie) Insert(pattern string) {
	parts := splitPath(pattern)

	t.root.insert(parts, pattern)
}

func (t *Trie) Search(url string) (pattern string) {
	// log.Infof("Searching url: %s", url)
	parts := splitPath(url)

	n := t.root.search(parts)
	return n.pattern
}

func splitPath(pattern string) []string {
	strs := strings.Split(pattern, "/")
	parts := make([]string, 0)

	for _, part := range strs {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}
