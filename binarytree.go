package main

import "fmt"

// Given a binary tree, find its minimum depth
// A minimum depth is the number of nodes along the shortest path from the root to the nearest leaf node.

//        3
//      /   \
//    9      20
//          /  \
//         15  30

// Answer: 2 (because 9 is the nearest leaf node reachable from root node using shortest path 2 (3 - 9))
var level int

type queueItem struct {
	node  *TreeNode
	level int
}

func findMinimumDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	nodeItem := &queueItem{node: root, level: 1}

	queue := []queueItem{*nodeItem}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if &item != nil {
			tNode := item.node
			if tNode.Left == nil && tNode.Right == nil {
				return item.level
			}

			if tNode.Left != nil {
				queue = append(queue, queueItem{node: tNode.Left, level: item.level + 1})
			}
			if tNode.Right != nil {
				queue = append(queue, queueItem{node: tNode.Right, level: item.level + 1})
			}
		}
	}

	return 0
}

func getLeftView(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	var result []int
	leftView(root, 1, 0, result)

	return result
}

func leftView(root *TreeNode, level int, maxLevel int, result []int) {

	if maxLevel < level {
		result = append(result, root.Value)
		fmt.Println(result)
		maxLevel = level
	}
	if root.Left != nil {
		leftView(root.Left, level+1, maxLevel, result)
	}

	if root.Right != nil {
		leftView(root.Right, level+1, maxLevel, result)
	}
}
