package main

import (
	"fmt"
	"leetcode/common"
	"math"
	"reflect"
)

func traversalRight(root *common.TreeNode, arr *[]int) {
	if root == nil {
		*arr = append(*arr, math.MaxInt)
		return
	}

	*arr = append(*arr, root.Val)
	traversalRight(root.Right, arr)
	traversalRight(root.Left, arr)

}

func traversalLeft(root *common.TreeNode, arr *[]int) {
	if root == nil {
		*arr = append(*arr, math.MaxInt)
		return
	}

	*arr = append(*arr, root.Val)
	traversalLeft(root.Left, arr)
	traversalLeft(root.Right, arr)
}

func isSymmetric(root *common.TreeNode) bool {
	if root == nil {
		return true
	}

	left := make([]int, 0)
	right := make([]int, 0)
	traversalLeft(root.Left, &left)
	traversalRight(root.Right, &right)

	fmt.Println(left)
	fmt.Println(right)

	return reflect.DeepEqual(left, right)

}

func main() {
	test := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val: 2,
			Left: &common.TreeNode{
				Val:   3,
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
			Val: 2,
			Left: &common.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &common.TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(isSymmetric(test))
}
