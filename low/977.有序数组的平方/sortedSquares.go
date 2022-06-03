package main

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	i := len(result) - 1
	left := 0
	right := len(nums) - 1

	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare < rightSquare {
			result[i] = rightSquare
			right--
			i--
		}

		if leftSquare >= rightSquare {
			result[i] = leftSquare
			i--
			left++
		}
	}

	return result
}
