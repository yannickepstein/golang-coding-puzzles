package main

import (
	"container/list"
)

func traverse(root *TreeNode) (levelOrderTraversal [][]int) {
	levelQueue := list.New()
	var remainingNodes int
	var currentLevel []int
	if root != nil {
		levelQueue.PushBack(root)
		remainingNodes = 1
	}
	for levelQueue.Len() > 0 {
		if remainingNodes == 0 {
			levelOrderTraversal = append(levelOrderTraversal, currentLevel)
			currentLevel = nil
			remainingNodes = levelQueue.Len()
		}
		front := levelQueue.Front()
		node := front.Value.(*TreeNode)
		currentLevel = append(currentLevel, node.Val)
		if node.Left != nil {
			levelQueue.PushBack(node.Left)
		}
		if node.Right != nil {
			levelQueue.PushBack(node.Right)
		}
		remainingNodes--
		levelQueue.Remove(front)
	}
	if currentLevel != nil {
		levelOrderTraversal = append(levelOrderTraversal, currentLevel)
	}
	return levelOrderTraversal
}
