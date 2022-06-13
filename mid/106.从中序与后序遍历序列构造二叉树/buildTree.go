package main

import (
	"leetcode/common"
)

func buildTree(inorder []int, postorder []int) *common.TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	root := &common.TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  nil,
		Right: nil,
	}
	i := 0
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	l1 := inorder[:i]
	l2 := postorder[:i]
	r1 := inorder[i+1:]
	r2 := postorder[i : len(postorder)-1]
	root.Left = buildTree(l1, l2)
	root.Right = buildTree(r1, r2)
	return root
}

func main() {
	common.DumpTreePreOrder(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}
