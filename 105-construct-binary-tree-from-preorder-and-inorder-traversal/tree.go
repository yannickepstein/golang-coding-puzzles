package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree creates a tree from the given preorder and postorder
// traversal
// if preorder defines [root, left, right], then we can use inorder
// to find the length of each subtree through iteration
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 || len(inorder) <= 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	leftSize := 0
	for i := 0; inorder[i] != root.Val; i++ {
		leftSize++
	}
	if leftSize == 0 {
		root.Left = nil
	} else {
		root.Left = BuildTree(preorder[1:leftSize+1], inorder[0:leftSize])
	}
	if leftSize+1 >= len(preorder) {
		root.Right = nil
	} else {
		root.Right = BuildTree(preorder[leftSize+1:], inorder[leftSize+1:])
	}
	return root
}
