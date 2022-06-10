package main

import (
	"leetcode/common"
)

func traversal() {

}

func buildTree(preorder []int, inorder []int) *common.TreeNode {
	var root *common.TreeNode

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	//根节点
	root = &common.TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	//找到左右子树
	left := 1   //前序遍历中左子树的起点
	right := -1 //前序遍历中右子树的起点
	for i, v := range inorder {
		if v == root.Val {
			right = i
		}
	}
	l1 := preorder[left : right+1]
	r1 := inorder[:right]
	l2 := preorder[right+1:]
	r2 := inorder[right+1:]
	root.Left = buildTree(l1, r1)
	root.Right = buildTree(l2, r2)
	return root
}

func main() {
	common.DumpTreePreOrder(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
