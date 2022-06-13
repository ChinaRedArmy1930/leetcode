package main

import (
	"fmt"
	"leetcode/common"
	"math"
)

//TODO
func recoverTree(root *common.TreeNode) {
	recoverTreeEx(root)
}
func recoverTreeEx(root *common.TreeNode) int {
	if root == nil {
		return ^math.MaxInt64
	}
	left := recoverTreeEx(root.Left)
	if left > root.Val {
		tmp := root.Left.Val
		root.Left.Val = root.Val
		root.Val = tmp
	}

	right := recoverTreeEx(root.Right)

	if left > right {
		return left
	}

	return right
}

func main() {
	test := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val:  3,
			Left: nil,
			Right: &common.TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
	}

	test2 := &common.TreeNode{
		Val: 3,
		Left: &common.TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &common.TreeNode{
			Val: 4,
			Left: &common.TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	_ = test
	recoverTree(test)
	common.DumpTreeInOrder(test)
	fmt.Println()
	recoverTree(test2)
	common.DumpTreeInOrder(test2)
}
