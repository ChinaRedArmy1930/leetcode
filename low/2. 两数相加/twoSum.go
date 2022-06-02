package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	for k1, v := range nums {
		for k2, v2 := range nums {
			if k1 == k2 {
				continue
			}
			if target == v+v2 {
				result = []int{k1, k2}
				return result
			}
		}
	}
	return result
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
