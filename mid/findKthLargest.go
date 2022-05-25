package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	numsLen := len(nums)
	if numsLen == 0 || numsLen < k {
		return -1
	}

	for i := 0; i < len(nums); i++ {
		value := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if value < nums[j] {
				value = nums[j]
				nums[i], nums[j] = nums[j], nums[i]
			}
		}

		if i+1 == k {
			return nums[i]
		}
	}

	return -1
}

func main() {

	type TestStruct struct {
		Array  []int
		Target int
	}
	testArrays := []TestStruct{
		{
			Array:  []int{3, 2, 1, 5, 6, 4},
			Target: 2,
		},
		{
			Array:  []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			Target: 4,
		},
	}

	for _, array := range testArrays {
		fmt.Println(findKthLargest(array.Array, array.Target))
	}

}
