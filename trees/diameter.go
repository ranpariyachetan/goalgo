package trees

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func TestDiameterOfBinaryTree() {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	r := diameterOfBinaryTree(&root)

	fmt.Println(r)
}

var x int

func diameterOfBinaryTree(root *TreeNode) int {
	x = 0
	solve(root)
	return x
}

func solve(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := solve(root.Left)

	right := solve(root.Right)

	temp := 1 + maxInt(left, right)

	x = maxInt(x, left+right)

	return temp
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
