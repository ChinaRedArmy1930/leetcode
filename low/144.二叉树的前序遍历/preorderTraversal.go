package main

import "leetcode/common"

func preorderTraversal(root *common.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)

	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)
	ret = append(ret, root.Val)
	ret = append(ret, left...)
	ret = append(ret, right...)

	return ret
}
