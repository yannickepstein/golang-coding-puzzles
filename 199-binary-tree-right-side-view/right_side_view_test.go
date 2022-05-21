package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRightSideView(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	expected := []int{1, 3, 4}

	view := RightSideView(root)

	if diff := cmp.Diff(expected, view); diff != "" {
		t.Error(diff)
	}
}
