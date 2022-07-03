package main

import (
	"leetcode/common"
)

func leverOrder(root *common.TreeNode) []int {
	p := make([]*common.TreeNode, 0)
	ret := make([]int, 0)
	p = append(p, root)

	for len(p) == 0 {
		q := make([]*common.TreeNode, 0)

		for i := 0; i < len(p); i++ {
			node := p[i]
			ret = append(ret, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		p = q
	}

	return ret
}

func twoSumBSTs(root1 *common.TreeNode, root2 *common.TreeNode, target int) bool {
	r1 := leverOrder(root1)
	r2 := leverOrder(root2)

	for i := 0; i < len(r1); i++ {
		for j := 0; j < len(r2); j++ {
			if r1[i]+r2[j] == target {
				return true
			}
		}
	}

	return false
}

func main() {

}
