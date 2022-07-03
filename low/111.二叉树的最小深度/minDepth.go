package main

import (
	"../../common"
	"container/list"
	"fmt"
	"math"
)

func traversal(root *common.TreeNode, minDepth *int, depth *int) {
	if root == nil {
		return
	}

	*depth++

	if root.Left == nil && root.Right == nil {
		if *depth < *minDepth {
			*minDepth = *depth
		}
		*depth--
		return
	}

	traversal(root.Left, minDepth, depth)
	traversal(root.Right, minDepth, depth)
	*depth--
}

func minDepth(root *common.TreeNode) int {
	depth := 0
	min := math.MaxInt64
	traversal(root, &min, &depth)
	if min == math.MaxInt64 {
		return 0
	}
	return min
}

func minDepth2(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	nodeList := list.New()
	nodeList.PushBack(root)
	depth := 1

	for nodeList.Len() > 0 {
		l := nodeList.Len()
		for i := 0; i < l; i++ {
			node := nodeList.Front().Value.(*common.TreeNode)
			nodeList.Remove(nodeList.Front())
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				nodeList.PushBack(node.Left)
			}
			if node.Right != nil {
				nodeList.PushBack(node.Right)
			}
		}
		depth++
	}

	return depth
}

func minDepth3(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	nodeList := list.New()
	nodeList.PushBack(root)
	step := 1
	for nodeList.Len() > 0 {
		l := nodeList.Len()
		for i := 0; i < l; i++ {
			node := nodeList.Front().Value.(*common.TreeNode)
			nodeList.Remove(nodeList.Front())
			if node.Left == nil && node.Right == nil {
				return step
			}
			if node.Left != nil {
				nodeList.PushBack(node.Left)
			}

			if node.Right != nil {
				nodeList.PushBack(node.Right)
			}
		}
		step++
	}

	return step
}

func main() {
	test := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val: 2,
			Left: &common.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &common.TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &common.TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(minDepth(test))
	fmt.Println(minDepth2(test))
}
