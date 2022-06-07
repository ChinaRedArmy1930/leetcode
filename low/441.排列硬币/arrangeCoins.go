package main

func cacl1_n(n int) (ret int) {
	ret = (n * (n + 1)) / 2

	return
}

func arrangeCoins2(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		if sum == n {
			return i
		}
		if sum > n {
			return i - 1
		}
	}
	return -1
}

func arrangeCoins(n int) int {
	start := 1
	end := n

	ans := 0

	for start <= end {
		mid := start + (end-start)/2
		t := cacl1_n(mid)
		if t <= n {
			ans = mid
			start = mid + 1

		}

		if t > n {
			end = mid - 1
		}
	}

	return ans

}
