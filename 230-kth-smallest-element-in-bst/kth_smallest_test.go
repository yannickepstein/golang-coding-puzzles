package main

import "testing"

func TestKthSmallest(t *testing.T) {
	// [1, 2, 3, 4, 5] k = 3 -> expected = 3
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	exp := 3

	res, _ := KthSmallest(root, 3)

	if res != exp {
		t.Errorf("expected %d, but got %d", exp, res)
	}
}
