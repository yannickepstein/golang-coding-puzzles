package main

import "math"

func IsValidBst(root *TreeNode) bool {
	return hasRange(root, math.MinInt, math.MaxInt)
}

// all nodes must be in range [lowerBound, upperBound]
func hasRange(root *TreeNode, lowerBound int, upperBound int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lowerBound || root.Val >= upperBound {
		return false
	}
	return hasRange(root.Left, lowerBound, root.Val) && hasRange(root.Right, root.Val, upperBound)
}
