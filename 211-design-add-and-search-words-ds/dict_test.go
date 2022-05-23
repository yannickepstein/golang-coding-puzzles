package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAdd(t *testing.T) {
	word := "word"
	exp := WordDictionary{
		EndsWord: false,
		Children: map[rune]*WordDictionary{
			'w': {
				EndsWord: false,
				Children: map[rune]*WordDictionary{
					'o': {
						EndsWord: false,
						Children: map[rune]*WordDictionary{
							'r': {
								EndsWord: false,
								Children: map[rune]*WordDictionary{
									'd': {
										EndsWord: true,
										Children: map[rune]*WordDictionary{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	dict := NewWordDictionary()
	dict.AddWord(word)

	if diff := cmp.Diff(dict, exp); diff != "" {
		t.Error(diff)
	}
}

func TestSearch(t *testing.T) {
	wd := WordDictionary{
		EndsWord: false,
		Children: map[rune]*WordDictionary{
			'w': {
				EndsWord: false,
				Children: map[rune]*WordDictionary{
					'o': {
						EndsWord: false,
						Children: map[rune]*WordDictionary{
							'r': {
								EndsWord: false,
								Children: map[rune]*WordDictionary{
									'd': {
										EndsWord: true,
										Children: map[rune]*WordDictionary{},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	word := "wo.d"

	if !wd.Search(word) {
		t.Errorf("expected to find '%q' in dictionary", word)
	}
}
