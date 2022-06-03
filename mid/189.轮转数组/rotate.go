package main

//尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
//你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k <= 0 {
		return
	}
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}

	for i := 0; i < k/2; i++ {
		nums[i], nums[k-i-1] = nums[k-i-1], nums[i]
	}

	for i := 0; i < (len(nums)-k)/2; i++ {
		nums[i+k], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i+k]

	}
}
