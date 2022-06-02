package main

import "fmt"

func search(nums []int, target int) int {
	i := 0
	start := i
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			end = mid - 1
		}

		if target > nums[mid] {
			start = mid + 1
		}
	}

	return -1
}

func main() {
	type TestCase struct {
		nums   []int
		target int
	}

	var test = []TestCase{
		{[]int{12}, 12},
		{[]int{-1, 0, 3, 5, 9, 12}, 9},
		{[]int{-1, 0, 3, 5, 9, 12, 12}, 12},
		{[]int{-1, 0, 3, 5, 9, 12, 12}, 1},
	}
	for _, testCase := range test {
		fmt.Println(search(testCase.nums, testCase.target))
	}
}
