package common

import "fmt"

// DumpTreePreOrder 前序遍历二叉树 根左右
func DumpTreePreOrder(head *TreeNode) {
	if head == nil {
		fmt.Print("nil ")
		return
	}
	fmt.Print(head.Val)
	fmt.Print(" ")
	DumpTreePreOrder(head.Left)
	DumpTreePreOrder(head.Right)
}
