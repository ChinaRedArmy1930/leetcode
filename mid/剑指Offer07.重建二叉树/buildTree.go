package main

import "leetcode/common"

//todo
func buildTree(preorder []int, inorder []int) *common.TreeNode {
	var root *common.TreeNode

	i := 0
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Val = preorder[0]
	root.Left = buildTree(preorder[1:i-1], inorder[0:i-1])
	root.Right = buildTree(preorder[i+1:], inorder[i:])
	return root
}

func main() {
	common.DumpTreePreOrder(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
