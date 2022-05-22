package main

import "testing"

// the max path sum of a binary tree can be described as:
// max {
//   max {root.Val + max path sum starting at the left subtree, root.Val + max path sum starting at the right subtree},
//   max path sum starting at the left subtree,
//   max path sum starting at the right subtree
// }
// we observe that if we solve max path sum starting at the left/right subtree first
// we can easily answer all three cases without repeating any work
func TestMaxPathSum(t *testing.T) {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	exp := 17

	res := MaxPathSum(root)

	if res != exp {
		t.Errorf("expected %d but received %d", exp, res)
	}
}
