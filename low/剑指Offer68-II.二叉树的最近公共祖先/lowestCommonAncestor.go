package main

import (
	"fmt"
	"leetcode/common"
)

func traversal(root *common.TreeNode, p, q *common.TreeNode, depth *int, max *int, ans **common.TreeNode) []*common.TreeNode {
	if root == nil {
		return nil
	}
	*depth++

	ret := make([]*common.TreeNode, 0)
	left := traversal(root.Left, p, q, depth, max, ans)
	right := traversal(root.Right, p, q, depth, max, ans)

	count := 0
	if root == p || root == q {
		count += 1
	}

	for _, node := range left {
		if node == p || node == q {

			count += 1

		}
	}

	for _, node := range right {
		if node == p || node == q {

			count += 1
		}
	}

	//左右两边子树找到了 p q 说明是公共的节点
	if count >= 2 {
		if *max < *depth {
			*max = *depth
			*ans = root
		}
	}
	ret = append(ret, root)
	ret = append(ret, left...)
	ret = append(ret, right...)
	*depth--
	return ret
}

func lowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	depth := 0
	max := 0
	ans := new(common.TreeNode)
	traversal(root, p, q, &depth, &max, &ans)
	return ans
}

func main() {
	test1 := &common.TreeNode{
		Val: 3,
		Left: &common.TreeNode{
			Val: 5,
			Left: &common.TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &common.TreeNode{
				Val: 2,
				Left: &common.TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &common.TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &common.TreeNode{
			Val: 1,
			Left: &common.TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &common.TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}
	_ = test1

	//fmt.Println(lowestCommonAncestor(test1, test1.Left, test1.Right).Val)
	fmt.Println(lowestCommonAncestor(test1, test1.Left, test1.Left.Right.Right).Val)
}
