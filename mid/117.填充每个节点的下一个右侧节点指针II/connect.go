package main

import (
	"leetcode/common"
)

//官方答案
func connectLeetCode(root *common.Node) *common.Node {
	start := root
	for start != nil {
		var nextStart, last *common.Node
		handle := func(cur *common.Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart
	}
	return root
}

func connect(root *common.Node) *common.Node {
	if root == nil {
		return nil
	}
	q := []*common.Node{root}

	for j := 0; len(q) > 0; j++ {
		var p []*common.Node
		firstEachLine := new(common.Node)
		firstEachLine = nil
		for i := 0; i < len(q); i++ {
			if q[i].Left == nil && q[i].Right == nil {
				continue
			}

			//处理一个节点下的左右节点
			if q[i].Left != nil && q[i].Right != nil {
				q[i].Left.Next = q[i].Right
			}

			if firstEachLine != nil {
				if q[i].Left != nil {
					firstEachLine.Next = q[i].Left
					firstEachLine = nil
				} else if q[i].Right != nil {
					firstEachLine.Next = q[i].Right
					firstEachLine = nil
				}
			}

			/* 找到下一行第一个需要连接的节点  */
			if q[i].Right != nil && firstEachLine == nil {
				firstEachLine = q[i].Right
			}

			if q[i].Left != nil && firstEachLine == nil {
				firstEachLine = q[i].Left
			}

			if q[i].Left != nil {
				p = append(p, q[i].Left)
			}

			if q[i].Right != nil {
				p = append(p, q[i].Right)
			}
		}
		q = p
	}

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
			Right: nil,
			Next:  nil,
		},
		Right: &common.Node{
			Val:  3,
			Left: nil,
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
	test2 := &common.Node{
		Val: 1,
		Left: &common.Node{
			Val: 2,
			Left: &common.Node{
				Val: 4,
				Left: &common.Node{
					Val:   7,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
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
			Val:  3,
			Left: nil,
			Right: &common.Node{
				Val:  6,
				Left: nil,
				Right: &common.Node{
					Val:   8,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
				Next: nil,
			},
			Next: nil,
		},
		Next: nil,
	}
	test3 := &common.Node{
		Val: 1,
		Left: &common.Node{
			Val: 2,
			Left: &common.Node{
				Val: 4,
				Left: &common.Node{
					Val:   8,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
				Right: nil,
				Next:  nil,
			},
			Right: &common.Node{
				Val:  5,
				Left: nil,
				Right: &common.Node{
					Val:   9,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
				Next: nil,
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
				Val:  7,
				Left: nil,
				Right: &common.Node{
					Val:   16,
					Left:  nil,
					Right: nil,
					Next:  nil,
				},
				Next: nil,
			},
			Next: nil,
		},
		Next: nil,
	}
	test4 := &common.Node{
		Val: 0,
		Left: &common.Node{
			Val:  -4,
			Left: nil,
			Right: &common.Node{
				Val:  -1,
				Left: nil,
				Right: &common.Node{
					Val: 3,
					Left: &common.Node{
						Val:   -2,
						Left:  nil,
						Right: nil,
						Next:  nil,
					},
					Right: nil,
					Next:  nil,
				},
				Next: nil,
			},
			Next: nil,
		},
		Right: &common.Node{
			Val: -3,
			Left: &common.Node{
				Val:  8,
				Left: nil,
				Right: &common.Node{
					Val: -9,
					Left: &common.Node{
						Val:   4,
						Left:  nil,
						Right: nil,
						Next:  nil,
					},
					Right: nil,
					Next:  nil,
				},
				Next: nil,
			},
			Right: nil,
			Next:  nil,
		},
		Next: nil,
	}
	_ = test
	_ = test2
	_ = test3
	_ = test4
	//ret := connect(test)
	//
	//fmt.Println(ret.Left.Next.Val)
	//fmt.Println(ret.Left.Left.Next.Val)

	//ret2 := connect(test2)
	//
	//fmt.Println(ret2.Left.Left.Left.Next.Val)
	//ret3 := connect(test3)
	//
	//fmt.Println(ret3.Left.Left.Next.Val)
	//fmt.Println(ret3.Left.Left.Left.Next.Val)
	//fmt.Println(ret3.Left.Left.Left.Next.Next.Val)
	ret4 := connectLeetCode(test4)
	common.DumpNodeLevelOrderEx(ret4)
	//fmt.Println(ret4.Left.Right.Right.Left.Next.Val)
}
