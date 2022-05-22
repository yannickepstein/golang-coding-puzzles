package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBuildTree(t *testing.T) {
	expectedTree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
	}
	preorder := []int{3, 9, 5, 4, 2, 11, 1, 7}
	inorder := []int{9, 3, 2, 4, 5, 1, 11, 7}

	tree := BuildTree(preorder, inorder)

	if diff := cmp.Diff(expectedTree, tree); diff != "" {
		t.Errorf(diff)
	}
}
