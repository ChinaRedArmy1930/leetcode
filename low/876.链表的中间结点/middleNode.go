package main

import (
	"fmt"
	"leetcode/common"
)

func middleNode(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	}
	tmp1 := head
	tmp2 := head

	for tmp1.Next != nil && tmp2.Next != nil {
		tmp1 = tmp1.Next
		if tmp2.Next.Next == nil {
			break
		}
		tmp2 = tmp2.Next.Next
	}

	return tmp1
}

func main() {
	head := common.BuildList([]int{1, 2, 3, 4, 5})
	common.DumpList(head)
	fmt.Println()
	common.DumpList(middleNode(head))
	fmt.Println()
	head2 := common.BuildList([]int{1, 2, 3, 3, 4, 5})
	common.DumpList(head2)
	fmt.Println()
	common.DumpList(middleNode(head2))
}
