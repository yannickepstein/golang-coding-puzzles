package main

// Count counts the number of good nodes in a binary tree
// a node is good, if all nodes on the path from the root are <= the node's value
func Count(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return 1 + dfsCount(root.Left, root.Val) + dfsCount(root.Right, root.Val)
	}
}

func dfsCount(root *TreeNode, maxVal int) int {
	if root == nil {
		return 0
	}

	if root.Val >= maxVal {
		return 1 + dfsCount(root.Left, root.Val) + dfsCount(root.Right, root.Val)
	} else {
		return dfsCount(root.Left, maxVal) + dfsCount(root.Right, maxVal)
	}
}
