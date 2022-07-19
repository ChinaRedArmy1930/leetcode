package main

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		for l+1 < r && nums[l] == nums[l+1] {
			l++
		}

		for l < r-1 && nums[r] == nums[r-1] {
			r--
		}
		mid := (l + r) >> 1
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return nums[l]
}
