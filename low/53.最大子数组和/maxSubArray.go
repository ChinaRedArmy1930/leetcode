package main

import (
	"math"
	"sort"
)

func max(x ...int) int {
	sort.Ints(x)
	return x[len(x)-1]
}

func min(x ...int) int {
	sort.Ints(x)
	return x[0]
}

func dp(nums []int, index int, dpTable []int) int {
	if index == 0 {
		return nums[0]
	}

	tmp := ^math.MaxInt64
	maxN := nums[0]
	for i := 1; i < len(nums); i++ {
		tmp = max(nums[i], dpTable[i-1]+nums[i])
		dpTable = append(dpTable, tmp)
		if maxN < tmp {
			maxN = tmp
		}
	}

	return maxN
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	//dpTable[i] => 以num[i]结尾的最大连续子数组和
	dpTable := make([]int, 0)
	dpTable = append(dpTable, nums[0])
	return dp(nums, 1, dpTable)
}
