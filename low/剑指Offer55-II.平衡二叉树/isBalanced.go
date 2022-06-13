package main

import (
	"fmt"
	"leetcode/common"
	"math"
)

func height(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	l := height(root.Left)
	r := height(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}

func isBalanced(root *common.TreeNode) bool {
	if root == nil {
		return true
	}

	return math.Abs(float64(height(root.Left)-height(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
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
		Val: 1,
		Left: &common.TreeNode{
			Val: 2,
			Left: &common.TreeNode{
				Val: 3,
				Left: &common.TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: &common.TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &common.TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &common.TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	_ = test
	_ = test2
	//fmt.Println(isBalanced(test))
	//fmt.Println()
	fmt.Println(isBalanced(test2))
}
