package main

import "container/list"

func findWords(board [][]byte, words []string) []string {
	t := NewTrie()
	for _, word := range words {
		t.Add(word)
	}
	gridWords := []string{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			start := coordinate{
				X: i,
				Y: j,
			}
			foundWords := wordsStartingAt(board, t, start)
			gridWords = append(gridWords, foundWords...)
		}
	}
	return gridWords
}

type coordinate struct {
	X int
	Y int
}

func wordsStartingAt(board [][]byte, t *Trie, start coordinate) []string {
	explore := list.List{}
	explore.PushBack(
		searchState{
			Coordinate: start,
			T:          t,
			Word:       "",
		},
	)
	foundWords := []string{}
	for explore.Len() > 0 {
		front := explore.Front()
		state := front.Value.(searchState)
		explore.Remove(front)
		b := board[state.Coordinate.X][state.Coordinate.Y]
		c, ok := state.T.Match(b)
		if ok {
			if c.EndsWord {
				foundWords = append(foundWords, state.Word)
			}
			right, canMoveRight := moveRight(board, state.Coordinate)
			if canMoveRight {
				explore.PushBack(searchState{
					Coordinate: *right,
					T:          c,
					Word:       state.Word + string(b),
				})
			}
			down, canMoveDown := moveDown(board, state.Coordinate)
			if canMoveDown {
				explore.PushBack(searchState{
					Coordinate: *down,
					T:          c,
					Word:       state.Word + string(b),
				})
			}
		}
	}
	return foundWords
}

type searchState struct {
	Coordinate coordinate
	T          *Trie
	Word       string
}

func moveRight(board [][]byte, c coordinate) (*coordinate, bool) {
	if c.X+1 < len(board[0]) {
		return &coordinate{
			X: c.X + 1,
			Y: c.Y,
		}, true
	}
	return nil, false
}

func moveDown(board [][]byte, c coordinate) (*coordinate, bool) {
	if c.Y+1 < len(board) {
		return &coordinate{
			X: c.X,
			Y: c.Y + 1,
		}, true
	}
	return nil, false
}
