package common

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type ListNode struct {
	Val  int
	Next *ListNode
}
