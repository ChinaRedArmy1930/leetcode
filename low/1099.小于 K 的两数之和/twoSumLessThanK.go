package main

import (
	"fmt"
	"math"
	"sort"
)

func twoSumLessThanK(nums []int, k int) int {
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1

	max := ^math.MaxInt16

	for left < right {
		sum := nums[left] + nums[right]
		if sum < k && sum > max {
			max = sum
		}

		if sum < k {
			left++
		} else {
			right--
		}
	}

	if max == ^math.MaxInt16 {
		max = -1
	}

	return max
}

func main() {
	fmt.Println(twoSumLessThanK([]int{34, 23, 1, 24, 75, 33, 54, 8}, 60))
}
