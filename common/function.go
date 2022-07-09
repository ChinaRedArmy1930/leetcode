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

	DumpTreePreOrder(head.Left)
	DumpTreePreOrder(head.Right)
}

func DumpTreeLevelOrder(head *TreeNode) {
	if head == nil {
		return
	}

	p := []*TreeNode{head}

	for i := 0; len(p) > 0; i++ {
		var q []*TreeNode
		for j := 0; j < len(p); j++ {
			node := p[j]
			if node == nil {
				fmt.Print("null ")
				continue
			}
			fmt.Print(strconv.Itoa(node.Val) + " ")
			q = append(q, node.Left)
			q = append(q, node.Right)
		}
		fmt.Println()
		p = q
	}
}

func DumpTreeInOrder(head *TreeNode) {
	if head == nil {
		fmt.Print("nil ")
		return
	}

	DumpTreePreOrder(head.Left)
	fmt.Print(head.Val)
	fmt.Print(" ")
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
	fmt.Println()
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

func DumpNodeLevelOrder(root *Node) {
	if root == nil {
		return
	}
	p := []*Node{root}
	for i := 0; len(p) > 0; i++ {
		var q []*Node
		fmt.Printf("%d: ", i)
		for j := 0; j < len(p); j++ {
			node := p[j]
			fmt.Printf("%d ", node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		p = q
		fmt.Println()
	}
}

func DumpNodeLevelOrderEx(root *Node) {
	if root == nil {
		return
	}

	fmt.Printf("%d: ", 1)
	fmt.Print(root.Val)
	fmt.Println()
	p := []*Node{root}
	lineHead := new(Node)
	lineHead = nil
	for i := 0; len(p) > 0; i++ {
		var q []*Node
		fmt.Printf("%d: ", i+1)
		for j := 0; j < len(p); j++ {
			node := p[j]
			if node.Left == nil && node.Right == nil {
				continue
			}
			if lineHead == nil && node.Left != nil {
				lineHead = node.Left
			}
			if lineHead == nil && node.Right != nil {
				lineHead = node.Right
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		for lineHead != nil {
			fmt.Printf("%d ", lineHead.Val)
			lineHead = lineHead.Next
		}

		p = q

		fmt.Println()
	}
}
