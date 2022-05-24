package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func rev(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	prev = nil

	for curr != nil {
		tmp := curr.Next
		curr.Next = prev

		prev = curr
		curr = tmp
	}

	return prev

}

func rev2(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	if head.Next == nil {
		return []int{head.Val}
	}

	return append(rev2(head.Next), head.Val)
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prv *ListNode
	prv = nil

	for curr != nil {
		tmp := curr.Next
		curr.Next = prv
		prv = curr
		curr = tmp
	}

	return prv
}

func reverseList1(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	prev = nil

	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	return prev
}

func buildList() *ListNode {
	head := new(ListNode)
	head.Val = 0
	h := head
	for _, i := range []int{1, 2, 3, 4, 5} {
		tmp := &ListNode{
			Val:  i,
			Next: nil,
		}
		h.Next = tmp
		h = h.Next
	}

	return head
}

func debugList(list *ListNode) {
	debug := list
	for debug != nil {
		fmt.Println(debug.Val)
		debug = debug.Next
	}
}

func main() {
	head := buildList()

	//debugList(reverseList(head))
	ret := rev2(head)
	fmt.Println(ret)
}
