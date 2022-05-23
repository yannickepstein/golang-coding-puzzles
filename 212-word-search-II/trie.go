package main

type Trie struct {
	EndsWord bool
	Children map[byte]*Trie
}

func NewTrie() *Trie {
	return &Trie{
		EndsWord: false,
		Children: map[byte]*Trie{},
	}
}

func (t *Trie) Add(s string) {
	if len(s) == 0 {
		t.EndsWord = true
		return
	}
	first := s[0]
	c, ok := t.Children[first]
	if ok {
		c.Add(s[1:])
	} else {
		c := NewTrie()
		t.Children[first] = c
		c.Add(s[1:])
	}
}

func (t *Trie) Match(b byte) (*Trie, bool) {
	c, ok := t.Children[b]
	return c, ok
}
