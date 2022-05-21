package main

import "testing"

func TestIsValidBst(t *testing.T) {
	bstRoot := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 10,
			},
		},
	}
	if !IsValidBst(bstRoot) {
		t.Error("Expected is valid to be true")
	}

	nonBstRoot := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 10,
			},
		},
	}
	if IsValidBst(nonBstRoot) {
		t.Error("Expected is valid to be false")
	}
}
