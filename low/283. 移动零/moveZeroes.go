package main

import "fmt"

func moveZeroes2(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == 0 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i := 0

	count := 0

	for _, v := range nums {
		if v == 0 {
			count += 1
		}
	}

	if count == len(nums) {
		return
	}

	for {
		if i == len(nums)-count {
			break
		}

		if nums[i] == 0 {
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}

			continue
		}
		i++
	}

	for i := len(nums) - count; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	testArrays := [][]int{
		{0, 1, 0, 3, 12},
		{0},
		{0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	for _, array := range testArrays {
		moveZeroes2(array)
	}

	fmt.Println(testArrays)
}
