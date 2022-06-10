package main

import (
	"fmt"
	"leetcode/common"
	"math"
)

func traversal(root *common.TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := traversal(root.Left, max)
	if left < 0 {
		left = 0
	}
	right := traversal(root.Right, max)
	if right < 0 {
		right = 0
	}
	if *max < left+right+root.Val {
		*max = left + right + root.Val
	}

	if left > right {
		return left + root.Val
	}

	return right + root.Val
}

func maxPathSum(root *common.TreeNode) int {
	max := ^math.MaxInt64
	traversal(root, &max)
	return max
}

func main() {
	test := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &common.TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	test2 := &common.TreeNode{
		Val: -10,
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

	test3 := &common.TreeNode{
		Val:   -3,
		Left:  nil,
		Right: nil,
	}

	test4 := &common.TreeNode{
		Val: 2,
		Left: &common.TreeNode{
			Val:   -1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	_ = test
	_ = test2
	_ = test3
	_ = test4
	//fmt.Println(maxPathSum(test))
	//fmt.Println(maxPathSum(test2))
	//fmt.Println(maxPathSum(test3))
	fmt.Println(maxPathSum(test4))
}
