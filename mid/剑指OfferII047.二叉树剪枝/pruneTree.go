package main

import (
	"fmt"
	"leetcode/common"
)

func traversal(root *common.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := traversal(root.Left)
	right := traversal(root.Right)

	allZeroLeft := true
	for _, v := range left {
		if v != 0 {
			allZeroLeft = false
		}
	}

	if allZeroLeft {
		root.Left = nil
	}

	allZeroRight := true
	for _, v := range right {
		if v != 0 {
			allZeroRight = false
		}
	}

	if allZeroRight {
		root.Right = nil
	}

	var ret []int

	ret = append(ret, left...)
	ret = append(ret, right...)
	ret = append(ret, root.Val)

	return ret
}

func pruneTree(root *common.TreeNode) *common.TreeNode {
	dummayPoint := new(common.TreeNode)
	dummayPoint.Left = root
	traversal(dummayPoint)
	return dummayPoint.Left
}

func main() {
	test := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &common.TreeNode{
			Val: 0,
			Left: &common.TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &common.TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}

	test2 := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val: 0,
			Left: &common.TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &common.TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &common.TreeNode{
			Val: 1,
			Left: &common.TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &common.TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}

	test3 := &common.TreeNode{
		Val:  0,
		Left: nil,
		Right: &common.TreeNode{
			Val: 0,
			Left: &common.TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &common.TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
		},
	}

	test4 := &common.TreeNode{
		Val:  0,
		Left: nil,
		Right: &common.TreeNode{
			Val: 0,
			Left: &common.TreeNode{
				Val:  1,
				Left: nil,
				Right: &common.TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &common.TreeNode{
				Val:  1,
				Left: nil,
				Right: &common.TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	_ = test
	_ = test2
	_ = test3
	_ = test4
	common.DumpTreePreOrder(pruneTree(test))
	fmt.Println()
	common.DumpTreePreOrder(pruneTree(test2))
	fmt.Println()
	common.DumpTreePreOrder(pruneTree(test3))
	fmt.Println()
	common.DumpTreePreOrder(pruneTree(test4))
}
