package main

import (
	"fmt"
	"leetcode/common"
)

func traversal(root *common.TreeNode, target int, ans *bool, sum *int) {
	if root == nil {
		return
	}

	*sum += root.Val

	if *sum == target && root.Left == nil && root.Right == nil {
		*ans = true
		return
	}

	traversal(root.Left, target, ans, sum)
	traversal(root.Right, target, ans, sum)
	*sum -= root.Val
}

func hasPathSum(root *common.TreeNode, targetSum int) bool {

	ans := false
	sum := 0
	traversal(root, targetSum, &ans, &sum)
	return ans
}

func main() {
	test := &common.TreeNode{
		Val: 3,
		Left: &common.TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &common.TreeNode{
			Val: 20,
			Left: &common.TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &common.TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	test2 := &common.TreeNode{
		Val: 5,
		Left: &common.TreeNode{
			Val: 4,
			Left: &common.TreeNode{
				Val: 11,
				Left: &common.TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &common.TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &common.TreeNode{
			Val: 8,
			Left: &common.TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &common.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}

	_ = test
	_ = test2
	//fmt.Println(hasPathSum(test, 12))
	fmt.Println(hasPathSum(test2, 22))
}
