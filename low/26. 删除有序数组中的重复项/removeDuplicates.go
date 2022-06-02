package main

import "fmt"

func removeDuplicates2(nums []int) int {
	left := 0
	right := 0

	for {
		if left == len(nums) || right == len(nums) {
			break
		}
		if nums[left] == nums[right] {
			right++
		} else {
			left++
			nums[left] = nums[right]
		}
	}

	return left + 1
}

func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	resultLen := numsLen
	if numsLen == 0 {
		return -1
	}
	value := nums[0]
	for i := 1; i < numsLen; i++ {
		if nums[i] == value {
			resultLen = resultLen - 1
			numsLen = numsLen - 1
			for j := i; j < numsLen; j++ {
				nums[j] = nums[j+1]
			}
			i--
		}

		if nums[i] > value {
			value = nums[i]
		}
	}

	return resultLen
}

func main() {
	testNums := []int{1, 1, 2, 2, 3, 3, 3, 3, 3, 3, 3, 5, 6, 7, 8, 9, 9, 9}
	//testNums2 := []int{1, 1, 2}
	fmt.Println(removeDuplicates2(testNums))
	fmt.Println(testNums)
}
