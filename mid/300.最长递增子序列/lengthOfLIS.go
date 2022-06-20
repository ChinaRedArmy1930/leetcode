package main

import (
	"fmt"
	"math"
)

func dp(nums []int) int {
	dpTable := make([]int, 0)
	dpTable = append(dpTable, 1)
	max := ^math.MaxInt64
	for i := 1; i < len(nums); i++ {
		tmp := ^math.MaxInt64
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if tmp < dpTable[j]+1 {
					tmp = dpTable[j] + 1
				}
			}
		}
		if tmp == ^math.MaxInt64 {
			tmp = 1
		}

		if max < tmp {
			max = tmp
		}

		dpTable = append(dpTable, tmp)
	}

	if max == ^math.MaxInt64 {
		max = 1
	}

	fmt.Println(dpTable)

	return max
}

func lengthOfLIS(nums []int) int {
	return dp(nums)
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
	fmt.Println(lengthOfLIS([]int{0}))
}
