package main

import (
	"fmt"
	"leetcode/common"
)

func travseral(root *common.TreeNode, depthMap map[int][]int, depth *int) {
	if root == nil {
		return
	}

	*depth++

	if _, ok := depthMap[*depth]; ok {
		depthMap[*depth] = append(depthMap[*depth], root.Val)
	} else {
		tmp := make([]int, 0)
		tmp = append(tmp, root.Val)
		depthMap[*depth] = tmp
	}

	travseral(root.Left, depthMap, depth)
	travseral(root.Right, depthMap, depth)

	*depth--
}

func levelOrder(root *common.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	//K -> 层数, V -> 层数对应的数据
	depthMap := make(map[int][]int, 0)
	depth := 0
	travseral(root, depthMap, &depth)

	result := make([][]int, len(depthMap))

	for k, v := range depthMap {
		result[k-1] = v
	}

	return result
}

func levelOrder2(root *common.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := []*common.TreeNode{root}
	ret := make([][]int, 0)
	for i := 0; len(q) > 0; i++ {
		p := make([]*common.TreeNode, 0)
		tmp := make([]int, 0)
		for j := 0; j < len(q); j++ {
			node := q[j]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}

			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
		ret = append(ret, tmp)
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

	fmt.Println(levelOrder2(test))
}
