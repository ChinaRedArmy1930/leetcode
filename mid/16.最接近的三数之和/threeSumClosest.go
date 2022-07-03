package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	c := math.MaxFloat64
	var ret int
	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			}

			if math.Abs(float64(sum-target)) < c {
				c = math.Abs(float64(sum - target))
				ret = sum
			} else if sum < target {
				l++
			} else if sum > target {
				r--
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
}
