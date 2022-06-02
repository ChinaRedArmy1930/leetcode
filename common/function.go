package common

import (
	"fmt"
	"strconv"
)

// DumpTreePreOrder 前序遍历二叉树 根左右
func DumpTreePreOrder(head *TreeNode) {
	if head == nil {
		fmt.Print("nil ")
		return
	}
	fmt.Print(head.Val)
	fmt.Print(" ")
	//if head.Left == nil && head.Right == nil {
	//	return
	//}
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

//index代表array下标
func findLR(array []string, index int) (left *TreeNode, right *TreeNode) {
	l := 2*index + 1
	r := l + 1

	left = nil
	right = nil
	if l < len(array) {
		i, err := strconv.Atoi(array[l])
		if err != nil {
			left = nil
		} else {
			left = &TreeNode{
				Val:   i,
				Left:  nil,
				Right: nil,
			}
		}
	}

	if r < len(array) {
		i, err := strconv.Atoi(array[r])
		if err != nil {
			right = nil
		} else {
			right = &TreeNode{
				Val:   i,
				Left:  nil,
				Right: nil,
			}
		}
	}
	return
}
