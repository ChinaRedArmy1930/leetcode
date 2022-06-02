package main

import "fmt"

func main() {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}

func findRepeatNumber(nums []int) int {
	numsMap := map[int]bool{}
	for _, v := range nums {
		if _, ok := numsMap[v]; ok {
			return v
		}
		numsMap[v] = true
	}

	return -1
}
