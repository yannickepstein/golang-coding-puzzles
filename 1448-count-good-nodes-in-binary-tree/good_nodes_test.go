package main

import (
	"testing"
)

func TestCount(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: -1,
				},
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	expected := 3

	numGoodNodes := Count(root)

	if numGoodNodes != expected {
		t.Errorf("Expected %d but got %d", expected, numGoodNodes)
	}
}
