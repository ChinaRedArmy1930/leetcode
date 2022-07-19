package main

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		mid := (l + r) >> 1
		if checkInRight(nums, mid, target) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if nums[l] != target {
		return -1
	}

	return l
}

func checkInRight(q []int, mid, target int) bool {
	if q[0] <= q[mid] {
		if !(target >= q[0] && target <= q[mid]) {
			return true
		}
	} else {
		if target > q[mid] && target < q[0] {
			return true
		}
	}

	return false
}

func checkInRight1(q []int, mid, target int) bool {
	if q[0] <= q[mid] {
		if q[0] <= target && q[mid] >= target {
			return false
		} else {
			return true
		}
	}

	if q[0] > q[mid] {
		if target > q[mid] && target < q[0] {
			return true
		}
	}

	return false
}
