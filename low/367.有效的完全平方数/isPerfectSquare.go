package main

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	if num == 0 {
		return false
	}

	start := 1
	end := num
	mid := start
	for start <= end {
		mid = start + (end-start)/2
		if mid*mid == num {
			return true
		}

		if mid*mid > num {
			end = mid - 1
		}

		if mid*mid < num {
			start = mid + 1
		}
	}
	return false
}
