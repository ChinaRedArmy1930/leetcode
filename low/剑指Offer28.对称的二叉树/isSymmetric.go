package main

import (
	"fmt"
	"leetcode/common"
	"math"
	"reflect"
)

func traversalLeft(root *common.TreeNode, arrs *[]int) {
	if root == nil {
		return
	}

	*arrs = append(*arrs, root.Val)

	if root.Left != nil {
		traversalLeft(root.Left, arrs)
	} else {
		*arrs = append(*arrs, ^math.MaxInt64)
	}
	if root.Right != nil {
		traversalLeft(root.Right, arrs)
	} else {
		*arrs = append(*arrs, ^math.MaxInt64)
	}
}

func traversalRight(root *common.TreeNode, arrs *[]int) {
	if root == nil {
		return
	}

	*arrs = append(*arrs, root.Val)
	if root.Right != nil {
		traversalRight(root.Right, arrs)
	} else {
		*arrs = append(*arrs, ^math.MaxInt64)
	}
	if root.Left != nil {
		traversalRight(root.Left, arrs)
	} else {
		*arrs = append(*arrs, ^math.MaxInt64)
	}

}

func preOrderLeft(root *common.TreeNode) []int {
	arrs := make([]int, 0)
	traversalLeft(root, &arrs)
	return arrs
}

func preOrderRight(root *common.TreeNode) []int {
	arrs := make([]int, 0)
	traversalRight(root, &arrs)
	return arrs
}

func isSymmetric(root *common.TreeNode) bool {
	if root == nil {
		return true
	}

	left := preOrderLeft(root.Left)
	right := preOrderRight(root.Right)

	fmt.Println(left)
	fmt.Println(right)

	return reflect.DeepEqual(left, right)
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

	test3 := &common.TreeNode{
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

	test4 := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val:  2,
			Left: nil,
			Right: &common.TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &common.TreeNode{
			Val:  2,
			Left: nil,
			Right: &common.TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}

	_ = test
	_ = test2
	_ = test3
	_ = test4
	fmt.Println(isSymmetric(test))
	fmt.Println(isSymmetric(test2))
	fmt.Println(isSymmetric(test3))
	fmt.Println(isSymmetric(test4))
}
