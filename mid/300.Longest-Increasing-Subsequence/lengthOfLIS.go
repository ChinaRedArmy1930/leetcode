package main

import (
	"fmt"
	"sort"
)

func max(x ...int) int {
	sort.Ints(x)
	return x[len(x)-1]
}

func dp(nums []int) []int {
	dpTable := make([]int, len(nums))
	dpTable[0] = 1

	for i := 1; i < len(nums); i++ {
		v := -1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				if dpTable[j] > v {
					v = dpTable[j]
				}
			}
		}

		if v == -1 {
			dpTable[i] = 1
		} else {
			dpTable[i] = v + 1
		}
	}

	fmt.Println(dpTable)
	return dpTable
}

func lengthOfLIS(nums []int) int {
	t := dp(nums)
	return max(t...)
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}
