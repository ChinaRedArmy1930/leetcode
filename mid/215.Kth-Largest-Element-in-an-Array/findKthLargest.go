package main

import (
	"fmt"
)

func quickSort(nums []int, l, r, k int) int {
	if l >= r {
		return nums[l]
	}

	i := l - 1
	j := r + 1
	x := nums[(l+r)/2]
	for i < j {
		for ok := true; ok; ok = nums[i] > x {
			i++
		}
		for ok := true; ok; ok = nums[j] < x {
			j--
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	if j-l+1 >= k {
		return quickSort(nums, l, j, k)
	}

	return quickSort(nums, j+1, r, k-(j-l+1))
}

func findKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, k)
}

func main() {
	var m, k int

	fmt.Scanf("%d %d", &m, &k)
	nums := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	fmt.Println(findKthLargest(nums, k))
}
