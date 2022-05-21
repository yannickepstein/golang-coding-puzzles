package main

import "container/list"

// step down the tree level by level
// in each level pick the right most node and append it to the result list

func RightSideView(root *TreeNode) (view []int) {
	levelQueue := list.New()
	levelLen := 0
	if root != nil {
		levelQueue.PushBack(root)
		levelLen = 1
	}
	for levelQueue.Len() > 0 {
		front := levelQueue.Front()
		levelLen--
		node := front.Value.(*TreeNode)
		levelQueue.Remove(front)
		if node.Left != nil {
			levelQueue.PushBack(node.Left)
		}
		if node.Right != nil {
			levelQueue.PushBack(node.Right)
		}
		if levelLen == 0 {
			view = append(view, node.Val)
			levelLen = levelQueue.Len()
		}
	}
	return view
}
