package main

import (
	"fmt"
	"leetcode/common"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//把大的放前面
func insertionSortList2(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	curr := head.Next //代表当前需要排序的节点
	var sortedHead, sortedTail, prev *common.ListNode
	sortedHead = head //代表已排序好的链表头
	sortedTail = head //代表已排序好的链表尾
	for curr != nil {
		next := curr.Next
		comp := sortedHead //代表比较时候使用的指针
		prev = nil         //代表比较时需要用到的前一个指针
		for comp != curr {
			if comp.Val < curr.Val {
				curr.Next = comp
				if prev != nil {
					prev.Next = curr
				}
				if comp == sortedHead {
					sortedHead = curr
				}
				sortedTail.Next = next
				break
			}
			prev = comp
			comp = comp.Next
		}
		if comp == curr {
			sortedTail = curr
		}
		curr = next
	}
	return sortedHead
}

//把小的节点放在前面
func insertionSortList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}

	curr := head.Next //代表当前需要排序的指针
	var sortedTail, sortedHead *common.ListNode
	sortedTail = head //代表已排列好的链表的尾指针
	sortedHead = head //代表当前已排列好的链表的头指针

	var prev *common.ListNode

	for curr != nil {
		comp := sortedHead //用来比较后插入的头指针
		next := curr.Next
		prev = nil //比较时候的前一个节点
		for comp != curr {
			if curr.Val < comp.Val {
				sortedTail.Next = next

				if prev != nil {
					prev.Next = curr
				}
				curr.Next = comp

				//如果是插入最头部,则将curr置为head
				if comp == sortedHead {
					sortedHead = curr
				}

				curr = next
				break
			}
			prev = comp
			comp = comp.Next
		}

		//说明比到了最后都没找到比curr还要小的数字
		if comp == curr {
			sortedTail = curr
			curr = curr.Next
		}
	}

	return sortedHead
}

func main() {
	testLists := []*common.ListNode{
		common.BuildList([]int{1, 3, 0, 4}),
		common.BuildList([]int{1, 3, 2, 4, 0}),
		common.BuildList([]int{0}),
	}

	for _, v := range testLists {
		common.DumpList(insertionSortList2(v))
		fmt.Println()
	}

}
