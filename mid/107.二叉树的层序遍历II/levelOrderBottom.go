package main

import (
	"fmt"
	"leetcode/common"
)

func levelOrderBottom(root *common.TreeNode) [][]int {
	p := make([]*common.TreeNode, 0)
	ret := make([][]int, 0)

	p = append(p, root)
	for len(p) != 0 {
		q := make([]*common.TreeNode, 0)
		line := make([]int, 0)
		for i := 0; i < len(p); i++ {
			node := p[i]
			line = append(line, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}

		}

		p = q
		ret = append(ret, line)
	}

	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
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

	fmt.Println(levelOrderBottom(test))
}
