package main

import (
	"fmt"
	"sort"
)

func threeSumSmaller2(nums []int, target int) int {
	count := 0
	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum >= target {
				r--
			} else {
				count += r - l
				l++

			}
		}
	}

	return count
}

func threeSumSmaller(nums []int, target int) int {
	sort.Ints(nums)
	ret := 0
	for i := 0; i < len(nums)-2; i++ {
		for k := i + 1; k < len(nums); k++ {
			for z := k + 1; z < len(nums); z++ {
				sum := nums[k] + nums[i] + nums[z]
				if sum < target {
					ret++
				} else if sum >= target {
					break
				}
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(threeSumSmaller2([]int{-2, 0, 1, 3}, 2))
	fmt.Println(threeSumSmaller([]int{}, 0))
	fmt.Println(threeSumSmaller([]int{0}, 0))
	fmt.Println(threeSumSmaller([]int{1, 1, -2}, 4))
}
