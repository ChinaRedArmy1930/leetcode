package main

import "leetcode/common"

func traverse(root *common.TreeNode, depth *int, ans *int) {
	if root == nil {
		return
	}

	*depth++

	if *depth > *ans {
		*ans = *depth
	}

	traverse(root.Left, depth, ans)
	traverse(root.Right, depth, ans)
	*depth--
}

func maxDepth2(root *common.TreeNode) int {
	depth := 0
	ans := 0
	traverse(root, &depth, &ans)
	return ans
}

func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}
