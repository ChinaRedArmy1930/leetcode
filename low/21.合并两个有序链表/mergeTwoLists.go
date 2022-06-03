package main

import (
	"leetcode/common"
)

func mergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	var head *common.ListNode
	head = nil
	curr := head
	currList1 := list1
	currList2 := list2

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		curr = list2
		currList2 = currList2.Next
	} else {
		curr = list1
		currList1 = currList1.Next
	}
	head = curr
	for {

		if currList2 == nil || currList1 == nil {
			if currList1 == nil {
				curr.Next = currList2
				break
			}

			if currList2 == nil {
				curr.Next = currList1
				break
			}
		}

		if currList1.Val < currList2.Val {
			curr.Next = currList1
			curr = currList1
			currList1 = currList1.Next
		} else {
			curr.Next = currList2
			curr = currList2
			currList2 = currList2.Next
		}
	}

	return head
}

func main() {
	list1 := common.BuildList([]int{2})
	list2 := common.BuildList([]int{1})
	common.DumpList(mergeTwoLists(list1, list2))
}
