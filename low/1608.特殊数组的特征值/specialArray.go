package main

func checkEqual(nums []int, target int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			count++
		}
	}
	if count == target {
		return 0
	} else if count > target {
		return 1
	} else {
		return -1
	}
}

func specialArray(nums []int) int {
	start := 0
	end := len(nums)

	for start <= end {
		mid := start + (end-start)/2
		if checkEqual(nums, mid) == 0 {
			return mid
		}

		if checkEqual(nums, mid) == 1 {
			start = mid + 1
		}

		if checkEqual(nums, mid) == -1 {
			end = mid - 1
		}
	}

	return -1
}
