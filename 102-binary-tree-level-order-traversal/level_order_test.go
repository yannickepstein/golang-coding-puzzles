package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 17,
			},
		},
	}
	expected := [][]int{{3}, {9, 20}, {15, 17}}

	levelOrderTraversal := traverse(root)

	if diff := cmp.Diff(expected, levelOrderTraversal); diff != "" {
		t.Error(diff)
	}
}
