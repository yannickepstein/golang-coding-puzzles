package main

import "math"

// the max path sum can be described as
// max {
//   max path starting at the root,
//   max path skipping the root and starting somewhere in the left subtree,
//   max path skipping the root and starting somewhere in the right subtree
// }
//
// max path starting at the root is defined as follows:
// root.Val + continuePath(root.Left) + continuePath(root.Right)
// where continuePath(root) is defined as the max of
// 1. end -> return 0
// 2. continue left -> return root.Val + continuePath(root.Left)
// 3. continue right -> return root.Val + continuePath(root.Right)
//
// We observe that continuePath(root) will not change for any node in the tree, regardless of
// where exactly its path originated from
// thus we will call continuePath(node) for every node that is not a direct descendent of the tree's root
// at least twice
// thus we should cache the results for each node
// this way an algorithm implementing the above logic, will only have to traverse the tree twice:
// - once to explore all possible starts of a path
// - once to explore all continuations of any path starting above it

func MaxPathSum(root *TreeNode) int {
	cache := map[*TreeNode]int{}
	return maxPathSum(root, cache)
}

func maxPathSum(root *TreeNode, cache map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	maxPathStartingLeft := math.MinInt
	maxPathStartingRight := math.MinInt
	if root.Left != nil {
		maxPathStartingLeft = maxPathSum(root.Left, cache)
	}
	if root.Right != nil {
		maxPathStartingRight = maxPathSum(root.Right, cache)
	}
	startAtRoot := root.Val + continuePath(root.Left, cache) + continuePath(root.Right, cache)
	return max(startAtRoot, max(maxPathStartingLeft, maxPathStartingRight))
}

func continuePath(root *TreeNode, cache map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	cachedVal, ok := cache[root]
	if ok {
		return cachedVal
	}
	cache[root] = max(max(root.Val+continuePath(root.Left, cache), root.Val+continuePath(root.Right, cache)), 0)
	return cache[root]
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
