package main

import (
	"fmt"
	"leetcode/common"
	"strconv"
)

func tree2str(root *common.TreeNode) string {
	if root == nil {
		return "#,"
	}

	result := strconv.Itoa(root.Val) + ","
	if root.Left == nil && root.Right == nil {
		return result
	}
	result += tree2str(root.Left)
	result += tree2str(root.Right)

	return result
}

func Tree(root *common.TreeNode, m map[string]bool) []*common.TreeNode {
	var result []*common.TreeNode
	if root == nil {
		return nil
	}

	rootStr := tree2str(root)
	if v, ok := m[rootStr]; ok {
		if v == false {
			result = append(result, root)
		}

		m[rootStr] = true

	} else {
		m[rootStr] = false
	}

	if root.Left != nil {
		result = append(result, Tree(root.Left, m)...)
	}

	if root.Right != nil {
		result = append(result, Tree(root.Right, m)...)
	}

	return result
}

func findDuplicateSubtrees(root *common.TreeNode) []*common.TreeNode {
	var nodeMap = make(map[string]bool, 0)
	var result = make([]*common.TreeNode, 0)
	tree2str2(root, nodeMap, &result)
	return result
}

func tree2str2(root *common.TreeNode, m map[string]bool, r *[]*common.TreeNode) string {
	if root == nil {
		return "#,"
	}

	rootStr := strconv.Itoa(root.Val) + ","
	leftStr := tree2str2(root.Left, m, r)
	rightStr := tree2str2(root.Right, m, r)
	result := rootStr + leftStr + rightStr

	if v, ok := m[result]; ok {
		if v == false {
			*r = append(*r, root)
		}

		m[result] = true

	} else {
		m[result] = false
	}

	fmt.Println(result)
	return result
}

func main() {
	//root := &common.TreeNode{
	//	Val: -64,
	//	Left: &common.TreeNode{
	//		Val: 12,
	//		Left: &common.TreeNode{
	//			Val:  -4,
	//			Left: nil,
	//			Right: &common.TreeNode{
	//				Val:   51,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//		Right: &common.TreeNode{
	//			Val:   -53,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &common.TreeNode{
	//		Val:  18,
	//		Left: nil,
	//		Right: &common.TreeNode{
	//			Val: 16,
	//			Left: &common.TreeNode{
	//				Val:   -93,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: &common.TreeNode{
	//				Val:   3,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//	},
	//}
	root2 := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val: 2,
			Left: &common.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &common.TreeNode{
			Val: 3,
			Left: &common.TreeNode{
				Val: 2,
				Left: &common.TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &common.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}
	//fmt.Println(findDuplicateSubtrees(root))
	fmt.Println(findDuplicateSubtrees(root2))

}
