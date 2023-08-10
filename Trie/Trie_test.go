package trie

import (
	"testing"
)

func TestTrie1(t *testing.T) {
	// patterns
	patterns := []string{
		"/",
		"/api",
		"/api/v1",
		"/hello",
		"/hello/world",
		"/:lang/doc",
		"/:name/*file",
	}
	// insert several patterns
	trie := NewTrie()
	for _, pattern := range patterns {
		trie.Insert(pattern)
	}
	// search url
	urls := []string{
		"/",
		"/api",
		"/api/v1",
		"/hello",
		"/hello/world",
		"/golang/doc",
		"/jack/anime/photos/genshinimpact/keqing.jpg",
	}
	for i, url := range urls {
		res := trie.Search(url)
		if res != patterns[i] {
			t.Errorf("Search url %s failed, res is %s", url, res)
		}
	}
}
