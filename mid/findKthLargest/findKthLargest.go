package main

import "fmt"

func Heap(nums []int, start int, l int) {
	left := 2*start + 1
	right := left + 1

	largest := start

	if left < l && nums[left] > nums[largest] {
		largest = left
	}

	if right < l && nums[right] > nums[largest] {
		largest = right
	}

	if start != largest {
		nums[largest], nums[start] = nums[start], nums[largest]
		Heap(nums, largest, l)
	}
}

//使用堆排序
func findKthLargest(nums []int, k int) int {
	if k-1 > len(nums) {
		return -1
	}
	buildHeap := func(nums []int) {
		for i := (len(nums) / 2) - 1; i >= 0; i-- {
			Heap(nums, i, len(nums))
		}
	}

	buildHeap(nums)

	for i := 0; i < k; i++ {
		nums[0], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[0]
		Heap(nums, 0, len(nums)-i-1)
	}

	fmt.Println(nums)

	return nums[len(nums)-k]
}

func main() {
	//arrays := []int{1, 4, 3, 5, 6, 3, 2, 7, 6, 9, 10, 8}
	arrays2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}

	fmt.Println(findKthLargest(arrays2, 4))
}
