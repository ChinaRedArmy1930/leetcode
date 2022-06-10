package main

func countOneLine(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] < 0 {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return len(nums) - start
}

func countNegatives(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		count += countOneLine(grid[i])
	}
	return count
}
