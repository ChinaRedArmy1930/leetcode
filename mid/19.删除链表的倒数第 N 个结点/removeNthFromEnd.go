package main

import (
	"fmt"
	"leetcode/common"
)

//定义一个虚拟节点方法
func removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	metaPoint := &common.ListNode{
		Val:  0,
		Next: head,
	}

	tmp1 := metaPoint
	tmp2 := metaPoint

	count := 0
	//将tmp2前进n个节点
	for tmp2.Next != nil && count < n {
		tmp2 = tmp2.Next
		count++
	}

	for tmp1.Next != nil && tmp2.Next != nil {
		tmp1 = tmp1.Next
		tmp2 = tmp2.Next
	}

	//删除tmp1.Next
	tmp1.Next = tmp1.Next.Next

	return metaPoint.Next
}

func removeNthFromEnd2(head *common.ListNode, n int) *common.ListNode {
	if head == nil {
		return nil
	}
	tmp1 := head
	tmp2 := head
	count := 0
	len := 0
	for tmp2.Next != nil && count < n {
		tmp2 = tmp2.Next
		count++
		len++
	}

	for tmp1.Next != nil && tmp2.Next != nil {
		tmp1 = tmp1.Next
		tmp2 = tmp2.Next
		len++
	}
	len = len + 1

	if len == 1 {
		if n != len {
			return head
		} else {
			return nil
		}
	}

	//如果是正数第一个节点
	if len == n {
		return head.Next
	}

	//删除tmp1.Next
	tmp1.Next = tmp1.Next.Next

	return head
}

func main() {
	/**/
	head := common.BuildList([]int{1, 2, 3, 4, 5, 6})
	tmp := removeNthFromEnd(head, 2)
	common.DumpList(tmp)
	fmt.Println()

	head = common.BuildList([]int{1})
	fmt.Println()
	tmp = removeNthFromEnd(head, 1)
	common.DumpList(tmp)
	fmt.Println()

	head2 := common.BuildList([]int{1, 2})
	common.DumpList(head2)
	fmt.Println()
	tmp2 := removeNthFromEnd(head2, 1)
	common.DumpList(tmp2)
	fmt.Println()

	head3 := common.BuildList([]int{1})
	common.DumpList(head3)
	fmt.Println()
	tmp3 := removeNthFromEnd(head3, 1)
	common.DumpList(tmp3)
	fmt.Println()

	head4 := common.BuildList([]int{1, 2})
	common.DumpList(head4)
	fmt.Println()
	tmp4 := removeNthFromEnd(head4, 2)
	common.DumpList(tmp4)
	fmt.Println()

}
