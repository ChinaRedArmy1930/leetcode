package main

import "fmt"

func guess(num int) int {
	target := 4
	if num > target {
		return -1
	}

	if num < target {
		return 1
	}

	return 0
}

func guessNumber(n int) int {
	start := 1
	end := n
	mid := 1
	for start < end {
		mid = start + (end-start)/2
		switch guess(mid) {
		case -1:
			end = mid - 1
		case 1:
			start = mid + 1
		case 0:
			return mid
		}
	}

	return start
}

func main() {
	fmt.Println(guessNumber(10))
}
