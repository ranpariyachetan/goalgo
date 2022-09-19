package trees

import (
	"fmt"
)

var res [][]int

func TestPathSumII() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	pathSum(root, 22)

	fmt.Println(res)
}

func pathSum(root *TreeNode, targetSum int) [][]int {

	arr := make([]int, 0)
	res = make([][]int, 0)

	getPathSum(root, targetSum, arr)

	return res
}

func getPathSum(root *TreeNode, targetSum int, arr []int) {

	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			newarr := append(arr, root.Val)
			res = append(res, newarr)
		}
		return
	}

	if root.Left != nil {
		leftarr := append(arr, root.Val)
		getPathSum(root.Left, targetSum-root.Val, leftarr)
	}

	if root.Right != nil {
		rightarr := append(arr, root.Val)
		getPathSum(root.Right, targetSum-root.Val, rightarr)
	}
}
