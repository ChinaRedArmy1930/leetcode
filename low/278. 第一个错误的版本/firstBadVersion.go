package main

import (
	"fmt"
)

func isBadVersion(n int) bool {
	return n == 2
}

func firstBadVersion(n int) int {
	start := 1
	end := n
	mid := 1
	for start < end {
		mid = start + (end-start)/2
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}

func main() {
	fmt.Println(firstBadVersion(3))
}
