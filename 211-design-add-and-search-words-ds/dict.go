package main

import "unicode/utf8"

type WordDictionary struct {
	EndsWord bool
	Children map[rune]*WordDictionary
}

func NewWordDictionary() WordDictionary {
	return WordDictionary{
		EndsWord: false,
		Children: map[rune]*WordDictionary{},
	}
}

func (wd *WordDictionary) AddWord(word string) {
	if word == "" {
		wd.EndsWord = true
		return
	}
	first, size := utf8.DecodeRuneInString(word)
	c, ok := wd.Children[first]
	if ok {
		c.AddWord(word[size:])
	} else {
		c := NewWordDictionary()
		wd.Children[first] = &c
		c.AddWord(word[size:])
	}
}

func (wd *WordDictionary) Search(word string) bool {
	if word == "" {
		return wd.EndsWord
	}
	first, size := utf8.DecodeRuneInString(word)
	c, ok := wd.Children[first]
	if ok {
		return c.Search(word[size:])
	} else if first == '.' {
		anyMatch := false
		for _, c := range wd.Children {
			anyMatch = anyMatch || c.Search(word[size:])
		}
		return anyMatch
	}
	return false
}
