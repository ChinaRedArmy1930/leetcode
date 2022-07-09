package main

import (
	"leetcode/common"
)

func rev(root *common.ListNode) *common.ListNode {
	if root == nil || root.Next == nil {
		return root
	}

	last := rev(root.Next)
	root.Next.Next = root
	root.Next = nil

	return last
}

func reverseBetween(head *common.ListNode, left int, right int) *common.ListNode {
	d := &common.ListNode{
		Val:  0,
		Next: head,
	}

	h := d

	tmp := head

	var leftNode *common.ListNode
	var rightNode *common.ListNode
	var prevNode *common.ListNode
	var nextNode *common.ListNode
	foundLeft := false
	foundRight := false
	count := 1
	_ = rightNode
	for tmp != nil {
		if foundLeft && foundRight {
			break
		}

		if count == left {
			leftNode = d.Next
			d.Next = nil
			prevNode = d
			foundLeft = true
		}

		if count == right {
			rightNode = tmp
			nextNode = tmp.Next
			tmp.Next = nil
			foundRight = true
		}

		count++
		d = tmp
		tmp = tmp.Next
	}

	last := rev(leftNode)
	prevNode.Next = last
	leftNode.Next = nextNode
	return h.Next
}

func main() {
	common.DumpList(reverseBetween(common.BuildList([]int{1, 2, 3, 4, 5}), 2, 4))
	common.DumpList(reverseBetween(common.BuildList([]int{5}), 1, 1))
	common.DumpList(reverseBetween(common.BuildList([]int{3, 5}), 1, 2))
}
