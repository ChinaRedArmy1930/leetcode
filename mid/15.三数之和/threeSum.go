package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, start int, target int) [][]int {
	left := start
	right := len(nums) - 1
	ret := make([][]int, 0)
	for left < right {
		sum := nums[left] + nums[right]
		l := nums[left]
		r := nums[right]
		if sum == target {
			ret = append(ret, []int{nums[left], nums[right]})
			left++
			right--
			for left < right && nums[left] == l {
				left++
			}

			for left < right && nums[right] == r {
				right--
			}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}

	return ret
}

func threeSum2(nums []int) (ret [][]int) {
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			left := nums[l]
			right := nums[r]
			if sum == 0 {
				ret = append(ret, []int{nums[l], nums[r], nums[i]})
				l++
				r--
				for l < r && nums[l] == left {
					l++
				}
				for l < r && nums[r] == right {
					r--
				}
			} else if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			}
		}

		for i < len(nums)-3 && nums[i] == nums[i+1] {
			i++
		}
	}

	return ret
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	ret := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		t := twoSum(nums, i+1, 0-nums[i])
		for k, _ := range t {
			t[k] = append(t[k], nums[i])
		}
		ret = append(ret, t...)
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return ret
}

func main() {
	/*
		for _, v := range threeSum([]int{-1, 0, 1, 2, -1, -4}) {
			fmt.Println(v)
		}
		fmt.Println("===========")
		for _, v := range threeSum([]int{-2, 0, 0, 2, 2}) {
			fmt.Println(v)
		}
		fmt.Println("===========")

	*/
	for _, v := range threeSum2([]int{-1, 0, 1, 2, -1, -4}) {
		fmt.Println(v)
	}
	fmt.Println("===========")
	for _, v := range threeSum2([]int{-2, 0, 0, 2, 2}) {
		fmt.Println(v)
	}
}
