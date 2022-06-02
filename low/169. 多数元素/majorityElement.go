package main

import "fmt"

func majorityElement(nums []int) int {
	var result int
	count := 0 //代表result的个数
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			result = nums[i]
			count++
			continue
		}

		if result != nums[i] { //不相等的话, 则算是抵消了一个result
			count--
		} else {
			count++
		}
	}

	return result
}

func main() {
	fmt.Println(majorityElement([]int{3, 3, 4}))
}
