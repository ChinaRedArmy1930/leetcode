package main

func searchOneLine(nums []int, target int) bool {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > target {
			end = mid - 1
		}

		if nums[mid] < target {
			start = mid + 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if searchOneLine(matrix[i], target) {
			return true
		}
	}
	return false
}
