package main

import (
	"fmt"
	"leetcode/common"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *common.TreeNode) string {
	if root == nil {
		return ""
	}

	ret := ""

	p := []*common.TreeNode{root}

	for i := 0; len(p) > 0; i++ {
		var q []*common.TreeNode
		for j := 0; j < len(p); j++ {
			node := p[j]

			if node == nil {
				ret += "null "
				continue
			}
			ret += strconv.Itoa(node.Val) + " "
			q = append(q, node.Left)
			q = append(q, node.Right)
		}
		ret += "|" //每一层加个标记
		p = q
	}

	return ret
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *common.TreeNode {
	if data == "" {
		return nil
	}

	level := strings.Split(data, "|")
	rootVal, _ := strconv.Atoi(strings.Split(strings.TrimSpace(level[0]), " ")[0])
	root := &common.TreeNode{Val: rootVal}
	p := []*common.TreeNode{root} //代表已处理的节点

	for i := 1; i < len(level); i++ {
		var q []*common.TreeNode
		pIndex := 0
		arr := strings.Split(strings.TrimSpace(level[i]), " ")
		for j := 0; j < len(arr); j += 2 {
			if arr[j] == "" {
				break
			}
			if arr[j] != "null" {
				num, _ := strconv.Atoi(arr[j])
				p[pIndex].Left = &common.TreeNode{Val: num}
				q = append(q, p[pIndex].Left)
			}

			if arr[j+1] != "null" {
				num, _ := strconv.Atoi(arr[j+1])
				p[pIndex].Right = &common.TreeNode{Val: num}
				q = append(q, p[pIndex].Right)
			}
			pIndex++
		}

		p = q
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

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
	test2 := &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &common.TreeNode{
			Val: 3,
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
	}
	test3 := &common.TreeNode{
		Val: 4,
		Left: &common.TreeNode{
			Val:   -7,
			Left:  nil,
			Right: nil,
		},
		Right: &common.TreeNode{
			Val: -3,
			Left: &common.TreeNode{
				Val: -9,
				Left: &common.TreeNode{
					Val: 9,
					Left: &common.TreeNode{
						Val: 6,
						Left: &common.TreeNode{
							Val:   0,
							Left:  nil,
							Right: nil,
						},
						Right: &common.TreeNode{
							Val:  6,
							Left: nil,
							Right: &common.TreeNode{
								Val: -1,
								Left: &common.TreeNode{
									Val:   -2,
									Left:  nil,
									Right: nil,
								},
								Right: nil,
							},
						},
					},
					Right: nil,
				},
				Right: &common.TreeNode{
					Val: -7,
					Left: &common.TreeNode{
						Val: -6,
						Left: &common.TreeNode{
							Val: 5,
							Left: &common.TreeNode{
								Val:   -4,
								Left:  nil,
								Right: nil,
							},
							Right: nil,
						},
						Right: nil,
					},
					Right: &common.TreeNode{
						Val: -6,
						Left: &common.TreeNode{
							Val:   9,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
			},
			Right: &common.TreeNode{
				Val: -3,
				Left: &common.TreeNode{
					Val:   -4,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
	}
	test4 := &common.TreeNode{
		Val: 3,
		Left: &common.TreeNode{
			Val: 2,
			Left: &common.TreeNode{
				Val:   3,
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
	}
	test5 := &common.TreeNode{
		Val: -1,
		Left: &common.TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		Right: &common.TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	_ = test1
	_ = test2
	_ = test3
	_ = test4
	_ = test5
	ser := Constructor()
	deser := Constructor()
	//data := ser.serialize(test1)
	//fmt.Println(data)
	//ans1 := deser.deserialize(data)
	//common.DumpTreeLevelOrder(ans1)

	//data2 := ser.serialize(test2)
	//fmt.Println(data2)
	//ans2 := deser.deserialize(data2)
	//common.DumpTreePreOrder(ans2)
	//data3 := ser.serialize(test3)
	//fmt.Println(data3)
	//ans3 := deser.deserialize(data3)
	//common.DumpTreePreOrder(ans3)
	//data4 := ser.serialize(test4)
	//fmt.Println(data4)
	//ans4 := deser.deserialize(data4)
	//common.DumpTreePreOrder(ans4)

	data5 := ser.serialize(test5)
	fmt.Println(data5)
	ans5 := deser.deserialize(data5)
	common.DumpTreeLevelOrder(ans5)
}
