package main

import "fmt"

func heap(nums []int, start int, l int) {
	left := start*2 + 1
	right := left + 1
	largest := start
	if left < l && nums[left] < nums[largest] {
		largest = left
	}

	if right < l && nums[right] < nums[largest] {
		largest = right
	}

	if start != largest {
		nums[start], nums[largest] = nums[largest], nums[start]
		heap(nums, largest, l)
	}

}

func getLeastNumbers(arr []int, k int) []int {
	buildHeap := func(nums []int) {
		for i := (len(nums) / 2) - 1; i >= 0; i-- {
			heap(arr, i, len(nums))
		}
	}

	buildHeap(arr)

	for i := 0; i < k; i++ {
		arr[0], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[0]
		heap(arr, 0, len(arr)-i-1)
	}

	fmt.Println(arr)

	return arr[len(arr)-k:]
}

func main() {
	fmt.Println(getLeastNumbers([]int{0, 1, 2, 3, 4, 0, 3}, 5))
}
