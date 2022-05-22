package main

import (
	"fmt"
)

// KthSmallest finds the k-th smallest element
// To do this it builds up a slice that contains the
// elements of the BST in order and then returns the
// k-th element of this slice
func KthSmallest(root *TreeNode, k int) (int, error) {
	inorderTraversal := inorder(root)
	if len(inorderTraversal) < k {
		return -1, fmt.Errorf("the BST does not have %d elements", k)
	}
	return inorderTraversal[k-1], nil
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var inorderTraversal []int
	if root.Left != nil {
		inorderTraversal = append(inorderTraversal, inorder(root.Left)...)
	}
	inorderTraversal = append(inorderTraversal, root.Val)
	if root.Right != nil {
		inorderTraversal = append(inorderTraversal, inorder(root.Right)...)
	}
	return inorderTraversal
}
