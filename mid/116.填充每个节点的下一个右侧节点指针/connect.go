package main

import (
	"fmt"
	"leetcode/common"
)

func traversal(root *common.Node) {
	if root == nil {
		return
	}

	if root.Left != nil && root.Right != nil && root.Left.Next == nil {
		root.Left.Next = root.Right
	}

	leftArray := make([]*common.Node, 0)
	rightArray := make([]*common.Node, 0)
	if root.Left != nil && root.Right != nil {
		tmp := root.Left
		for tmp.Right != nil {
			rightArray = append(rightArray, tmp.Right)
			tmp = tmp.Right
		}

		tmp = root.Right
		for tmp.Left != nil {
			leftArray = append(leftArray, tmp.Left)
			tmp = tmp.Left
		}

		for i := 0; i < len(leftArray); i++ {
			if rightArray[i].Next == nil {
				rightArray[i].Next = leftArray[i]
			}

		}

	}

	traversal(root.Left)
	traversal(root.Right)
}

func connect(root *common.Node) *common.Node {
	traversal(root)
	return root
}

func main() {
	test := &common.Node{
		Val: 1,
		Left: &common.Node{
			Val: 2,
			Left: &common.Node{
				Val:   4,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &common.Node{
				Val:   5,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		},
		Right: &common.Node{
			Val: 3,
			Left: &common.Node{
				Val:   6,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &common.Node{
				Val:   7,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		},
		Next: nil,
	}

	ret := connect(test)

	fmt.Println(ret.Left.Next.Val)
	fmt.Println(ret.Left.Left.Next.Val)
	fmt.Println(ret.Left.Left.Next.Next.Val)
}
