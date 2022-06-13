package main

import (
	"fmt"
	"math"
)

type coinsNode struct {
	Val   int
	One   *coinsNode
	Three *coinsNode
	Five  *coinsNode
}

type nCoinsNode struct {
	Val   int
	Nodes []*nCoinsNode
}

// CreateNodes 生成N叉树
func CreateNodes(coins []int, amount int) *nCoinsNode {
	if amount < 0 {
		return nil
	}

	root := &nCoinsNode{Val: amount}

	for i := 0; i < len(coins); i++ {
		node := CreateNodes(coins, amount-coins[i])
		if node == nil {
			continue
		}
		root.Nodes = append(root.Nodes, node)
	}

	return root
}

//如果已经创建过
func CreateNodesEx(coins []int, amount int, nodeMap map[int]*nCoinsNode) *nCoinsNode {
	if amount < 0 {
		return nil
	}

	if v, ok := nodeMap[amount]; ok {
		return v
	}

	root := &nCoinsNode{Val: amount}

	for i := 0; i < len(coins); i++ {
		node := CreateNodes(coins, amount-coins[i])
		if node == nil {
			continue
		}
		root.Nodes = append(root.Nodes, node)
	}

	nodeMap[amount] = root

	return root
}

//优化: 获取每个节点的下的最小深度
func traversalNEx(root *nCoinsNode, depth *int, min *int, coinsMap map[int]int) int {
	if root == nil {
		return 0
	}
	*depth++

	if root.Val == 0 {
		if *min > *depth {
			*min = *depth
		}
		*depth--
		return 0
	}

	// v 代表节点下的最小深度
	if v, ok := coinsMap[root.Val]; ok {
		//如果节点遍历过, 并且当前深度+最小深度 比 min 还要小,则修改min
		if *min > v+*depth {
			*min = v + *depth
		}
		//不要往下遍历了
		*depth--
		return v
	}

	thisMin := math.MaxInt64
	for i := 0; i < len(root.Nodes); i++ {
		tmp := traversalNEx(root.Nodes[i], depth, min, coinsMap)
		if tmp < thisMin {
			thisMin = tmp
		}
	}

	coinsMap[root.Val] = thisMin + 1

	*depth--
	return thisMin + 1
}

func traversalN(root *nCoinsNode, depth *int, min *int) {
	if root == nil {
		return
	}

	*depth++

	//如果是叶子节点(子节点的个数为0), 且值为0,则说明符合条件的链条
	if len(root.Nodes) == 0 && root.Val == 0 && *min > *depth {
		*min = *depth
	}

	for i := 0; i < len(root.Nodes); i++ {
		traversalN(root.Nodes[i], depth, min)
	}

	*depth--
}

func CalcN(root *nCoinsNode) int {
	depth := 0
	min := math.MaxInt64
	coinsMap := make(map[int]int, 0)
	traversalNEx(root, &depth, &min, coinsMap)
	//fmt.Println(ret)
	if min == math.MaxInt64 {
		return -1
	}

	return min - 1
}

func traversal(coins []int, depth *int, min *int, amount int, coinsMap map[int]int) int {
	if amount <= 0 {
		return 0
	}
	*depth++
	if v, ok := coinsMap[amount]; ok {
		if v+*depth < *min {
			*min = v + *depth - 1
		}
		*depth--
		return v
	}

	thisMini := math.MaxInt64

	for i := 0; i < len(coins); i++ {
		if amount-coins[i] == 0 && *min > *depth {
			*min = *depth
		}

		if amount-coins[i] < 0 {
			continue
		}

		tmp := traversal(coins, depth, min, amount-coins[i], coinsMap)
		if thisMini > tmp {
			thisMini = tmp
		}
	}
	*depth--

	if thisMini == math.MaxInt64 {
		return math.MaxInt64
	}

	coinsMap[amount] = thisMini + 1
	return thisMini + 1
}

func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	depth := 0
	min := math.MaxInt64
	coinsMap := make(map[int]int, 0)
	traversal(coins, &depth, &min, amount, coinsMap)

	if min == math.MaxInt64 {
		return -1
	}

	return coinsMap[amount]
}

func coinChange(coins []int, amount int) int {
	nodeMap := make(map[int]*nCoinsNode, 0)
	root := CreateNodesEx(coins, amount, nodeMap)
	return CalcN(root)
}

func coinChangedp(coins []int, amount int) int {
	dp := make([]int, amount+1)
	max := math.MaxInt32
	for i := 0; i <= amount; i++ {
		dp[i] = max
	}
	dp[0] = 0

	for a := 1; a <= amount; a++ {
		for _, coin := range coins {
			if a < coin {
				continue
			}
			dp[a] = min(dp[a], dp[a-coin]+1)
		}
	}

	num := dp[amount]
	if num == max {
		return -1
	}
	return num
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	//for i := 1000; i < 2000; i++ {
	//	my := coinChange2([]int{186, 419, 83, 408}, i)
	//	ok := coinChangedp([]int{186, 419, 83, 408}, i)
	//
	//	if my != ok {
	//		fmt.Printf("i = %d my = %d, ok = %d \n", i, my, ok)
	//	} else {
	//		//fmt.Print(my)
	//
	//	}
	//}
	fmt.Println(coinChange2([]int{186, 419, 83, 408}, 4249))
	//fmt.Println(coinChangedp([]int{186, 419, 83, 408}, 4249))
	//fmt.Println(coinChange2([]int{1, 2, 5, 10}, 18))
	//fmt.Println(coinChange2([]int{2}, 3))
	//fmt.Println(coinChange2([]int{1, 3, 5}, 11))
	//for i := 0; i < 100; i++ {
	//	my := coinChange2([]int{1, 3, 5}, i)
	//	ok := coinChangedp([]int{1, 3, 5}, i)
	//
	//	if my != ok {
	//		fmt.Printf("i = %d my = %d, ok = %d \n", i, my, ok)
	//	} else {
	//		//fmt.Print(my)
	//		fmt.Printf(" i = %d ", i)
	//	}
	//}
}
