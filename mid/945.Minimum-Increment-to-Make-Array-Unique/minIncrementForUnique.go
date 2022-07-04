package main

import (
	"sort"
)

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			t := nums[i]
			nums[i] = nums[i-1] + 1
			count += nums[i] - t
		}
	}

	return count
}
