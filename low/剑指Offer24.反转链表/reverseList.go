package main

import (
	"fmt"
	"leetcode/common"
)

func rev(head *common.ListNode) *common.ListNode {
	curr := head
	var prev *common.ListNode
	prev = nil

	for curr != nil {
		tmp := curr.Next
		curr.Next = prev

		prev = curr
		curr = tmp
	}

	return prev

}

func rev2(head *common.ListNode) []int {
	if head == nil {
		return []int{}
	}

	if head.Next == nil {
		return []int{head.Val}
	}

	return append(rev2(head.Next), head.Val)
}

func reverseList(head *common.ListNode) *common.ListNode {
	curr := head
	var prv *common.ListNode
	prv = nil

	for curr != nil {
		tmp := curr.Next
		curr.Next = prv
		prv = curr
		curr = tmp
	}

	return prv
}

func reverseList1(head *common.ListNode) *common.ListNode {
	curr := head
	var prev *common.ListNode
	prev = nil

	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	return prev
}

func main() {
	head := common.BuildList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	common.DumpList(reverseList(head))
	ret := rev2(head)
	fmt.Println(ret)
}
