package main

func rotate(nums []int, k int) {
	numsLen := len(nums)
	realRotate := k % numsLen

	result := make([]int, 0)

	result = append(result, nums[numsLen-realRotate:]...)
	result = append(result, nums[:numsLen-realRotate]...)
	copy(nums, result)
}

func main() {
	testNums := []int{1, 2, 3, 4, 5, 6, 7}
	testNums2 := []int{-1, -100, 3, 99}
	rotate(testNums, 3)
	rotate(testNums2, 2)
}
