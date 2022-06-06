package main

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	start := 1
	end := x
	mid := start
	for start <= end {
		mid = start + (end-start)/2
		if mid*mid == x {
			break
		}

		if mid*mid < x && (mid+1)*(mid+1) > x {
			break
		}

		if mid*mid < x {
			start = mid + 1

		}

		if mid*mid > x {
			end = mid - 1
		}
	}

	return mid

}
