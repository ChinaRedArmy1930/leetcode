package main

import (
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}

	old := make([]int, len(nums))
	copy(old, nums)
	sort.Ints(nums)

	left := 0
	right := len(nums) - 1

	i := 0
	for nums[i] == old[i] {
		left++
		i++
	}

	i = len(nums) - 1
	for nums[i] == old[i] {
		i--
		right--
	}

	return right - left + 1
}
