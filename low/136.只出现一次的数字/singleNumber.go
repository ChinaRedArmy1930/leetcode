package main

import "fmt"

func singleNumber(nums []int) int {

	for i := 0; i < len(nums); i++ {
		value := nums[i]
		j := i - 1
		for j = i - 1; j >= 0; j-- {
			if nums[j] < value {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}

		nums[j+1] = value
	}

	fmt.Println(nums)

	j := 0
	for j < len(nums)-1 {
		if nums[j] != nums[j+1] {
			return nums[j]
		}
		j += 2
	}

	return nums[len(nums)-1]
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1, 0, 1}))
}
