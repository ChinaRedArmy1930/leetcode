package main

func traversal(nums []int, m *[]int, ret *[][]int) {
	if len(*m) == len(nums) {
		*ret = append(*ret, *m)
		return
	}

	for i := 0; i < len(nums); i++ {
		found := false

		for j := 0; j < len(*m); j++ {
			if (*m)[j] == nums[i] {
				found = true
				break
			}
		}

		if found {
			continue
		}
		c := make([]int, len(*m))
		copy(c, *m)
		*m = append(*m, nums[i])
		traversal(nums, m, ret)
		*m = make([]int, len(c))
		copy(*m, c)
	}
}

func permute(nums []int) [][]int {
	var ret [][]int
	var num []int
	traversal(nums, &num, &ret)
	return ret
}
