package main

import (
	"fmt"
	"leetcode/common"
)

func traversal(root *common.TreeNode, arrs *[]int) {
	if root == nil {
		return
	}

	*arrs = append(*arrs, root.Val)
	traversal(root.Left, arrs)
	traversal(root.Right, arrs)
}

func preorderTraversal(root *common.TreeNode) []int {
	arrs := make([]int, 0)
	traversal(root, &arrs)
	return arrs
}

func preorderTraversal1(root *common.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)

	left := preorderTraversal1(root.Left)
	right := preorderTraversal1(root.Right)
	ret = append(ret, root.Val)
	ret = append(ret, left...)
	ret = append(ret, right...)

	return ret
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

	fmt.Println(preorderTraversal(test))
}
