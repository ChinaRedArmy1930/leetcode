package main

func search(q []int, target int) bool {
	l := 0
	r := len(q) - 1

	for l < r {
		for l+1 < r && q[l] == q[l+1] {
			l++
		}
		for l < r-1 && q[r] == q[r-1] {
			r--
		}
		mid := (l + r + 1) >> 1

		if check(q, mid, target) {
			l = mid

		} else {
			r = mid - 1

		}
	}

	if q[l] == target {
		return true
	}

	return false
}

func check(q []int, mid, target int) bool {
	if q[0] < q[mid] {
		if !(target >= q[0] && target < q[mid]) {
			return true
		}
	} else {
		if target >= q[mid] && target < q[0] {
			return true
		}
	}

	return false
}
