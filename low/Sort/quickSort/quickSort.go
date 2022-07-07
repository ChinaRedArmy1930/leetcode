package main

import (
	"fmt"
	"math/rand"
)

func partition(nums []int, start, end int) int {
	index := end
	if end-start != 0 {
		index = start + rand.Intn(end-start)
	}
	nums[index], nums[end] = nums[end], nums[index]
	value := nums[end]
	j := start
	for i := start; i < end; i++ {
		if nums[i] < value {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	nums[j], nums[end] = nums[end], nums[j]
	return j
}

func quickR(nums []int, start, end int) {
	if start > end {
		return
	}
	mid := partition(nums, start, end)
	quickR(nums, start, mid-1)
	quickR(nums, mid+1, end)
}

func sortArray(nums []int) []int {
	quickR(nums, 0, len(nums)-1)
	return nums
}

func main() {
	fmt.Println(sortArray([]int{13, 5, 6, 7, 8, 3, 2, 9, 1, 0, 7, 5, 4, 3}))
}
