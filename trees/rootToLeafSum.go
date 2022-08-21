package trees

import "fmt"

// https://leetcode.com/problems/path-sum/submissions/

func TestHasPathSum() {
	var root *TreeNode

	r := hasPathSum(root, 0)

	fmt.Println(r)

	root = &TreeNode{}
	root.Val = 5
	root.Left = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 11,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 2},
		},
	}
	root.Right = &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 13,
		},
		Right: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	r = hasPathSum(root, 22)

	fmt.Println(r)
}

func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)

}
