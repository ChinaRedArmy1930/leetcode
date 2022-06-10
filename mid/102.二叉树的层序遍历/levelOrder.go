package main

import (
	"fmt"
	"leetcode/common"
)

func traversal(root *common.TreeNode, result map[int][]int, depth int) {
	if root == nil {
		return
	}

	depth++
	if root.Left != nil {
		result[depth] = append(result[depth], root.Left.Val)
	}

	if root.Right != nil {
		result[depth] = append(result[depth], root.Right.Val)
	}

	traversal(root.Left, result, depth)
	traversal(root.Right, result, depth)
	depth--
}
func levelOrder(root *common.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make(map[int][]int, 0)
	result[0] = []int{root.Val}
	traversal(root, result, 0)

	ret := make([][]int, 0)
	for i := 0; i < len(result); i++ {
		ret = append(ret, result[i])
	}

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

	fmt.Println(levelOrder(test))
}
