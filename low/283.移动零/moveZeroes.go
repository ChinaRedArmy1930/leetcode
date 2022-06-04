package main

func moveZeroes3(nums []int) {
	ZeroNum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[ZeroNum] = nums[i]
			ZeroNum++
		}
	}

	for i := ZeroNum; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroes2(nums []int) {
	l := 0
	r := 0
	for r < len(nums) {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}

func moveZeroes(nums []int) {
	i := 0 //第一个0
	j := 0 //第一个不为0

	for i < len(nums) && j < len(nums) {
		if nums[i] != 0 {
			i++
		}

		if nums[j] == 0 {
			j++
		}

		if i >= len(nums) || j >= len(nums) {
			break
		}

		if nums[i] == 0 && nums[j] != 0 {
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			} else {
				j++
			}
		}
	}
}
