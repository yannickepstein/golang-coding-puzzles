package trie

import "unicode/utf8"

type Trie struct {
	EndsWord bool
	Children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{
		EndsWord: false,
		Children: map[rune]*Trie{},
	}
}

func (t *Trie) Insert(word string) {
	if word == "" {
		t.EndsWord = true
		return
	}
	start, size := utf8.DecodeRuneInString(word)
	c, ok := t.Children[start]
	if ok {
		c.Insert(word[size:])
	} else {
		c := Constructor()
		t.Children[start] = &c
		c.Insert(word[size:])
	}
}

func (t *Trie) Search(word string) bool {
	if word == "" {
		return t.EndsWord
	}
	start, size := utf8.DecodeRuneInString(word)
	c, ok := t.Children[start]
	if !ok {
		return false
	}
	return c.Search(word[size:])
}

func (t *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return true
	}
	start, size := utf8.DecodeRuneInString(prefix)
	c, ok := t.Children[start]
	if !ok {
		return false
	}
	return c.StartsWith(prefix[size:])
}
