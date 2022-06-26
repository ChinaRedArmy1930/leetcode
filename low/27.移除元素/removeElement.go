package main

func removeElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums)-count; i++ {
		if nums[i] == val {
			for len(nums)-count > i {
				if nums[len(nums)-count-1] != val {
					nums[i], nums[len(nums)-count-1] = nums[len(nums)-count-1], nums[i]
					break
				} else {
					count++
				}
			}
		}
	}

	//	fmt.Println(nums)

	return len(nums) - count
}

func main() {
	removeElement([]int{3, 2, 2, 3}, 3)
	removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
}
