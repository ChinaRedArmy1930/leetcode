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

func BuildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := new(ListNode)
	head.Val = nums[0]
	h := head
	for _, i := range nums[1:] {
		tmp := &ListNode{
			Val:  i,
			Next: nil,
		}
		h.Next = tmp
		h = h.Next
	}

	return head
}

func DumpList(list *ListNode) {
	debug := list
	for debug != nil {
		fmt.Print(debug.Val)
		debug = debug.Next
	}
}
