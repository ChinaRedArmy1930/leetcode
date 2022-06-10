package main

import "leetcode/common"

func maxDepth(root *common.TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left, max)
	right := maxDepth(root.Right, max)

	if left+right > *max {
		*max = left + right
	}

	if left > right {
		return left + 1
	}

	return right + 1
}

func traversal(root *common.TreeNode) {
	if root == nil {
		return
	}
	traversal(root.Left)
	traversal(root.Right)
}

func diameterOfBinaryTree(root *common.TreeNode) int {
	max := 0
	maxDepth(root, &max)
	return max
}
