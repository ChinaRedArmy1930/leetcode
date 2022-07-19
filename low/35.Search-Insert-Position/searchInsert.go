package main

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if nums[l] == target {
		return l
	}

	if nums[l] < target {
		return l + 1
	} else {
		return l
	}
}
