package main

import (
	"fmt"
	"math"
	"sort"
)

func max(x ...int) int {
	sort.Ints(x)
	return x[len(x)-1]
}

func dp(nums []int) int {
	//1. dpTable[i] 代表以i结尾的最长递增子序列
	dpTable := make([]int, len(nums))
	//2. dpTable初始化
	for i := 0; i < len(dpTable); i++ {
		dpTable[i] = 1
	}
	//3.状态转移方程

	//4. 遍历顺序 从左到右

	for i := 1; i < len(nums); i++ {
		maxN := ^math.MaxInt16
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dpTable[j] > maxN {
					maxN = dpTable[j]
				}
			}
		}
		if maxN != ^math.MaxInt16 {
			dpTable[i] = maxN + 1
		}
	}

	//5. 打印数组
	fmt.Println(dpTable)
	return max(dpTable...)
}

func lengthOfLIS(nums []int) int {
	return dp(nums)
}
