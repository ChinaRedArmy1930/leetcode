package main

import "strconv"

func traversal(nums []int, m *[]int, numsMap *map[int]int, foundMap *map[string]bool, charMap *map[int]string, ret *[][]int) {
	if len(*m) == len(nums) {
		key := ""
		for _, v := range *m {
			key += (*charMap)[v]
		}

		if _, ok := (*foundMap)[key]; ok {
			return
		}

		*ret = append(*ret, *m)
		(*foundMap)[key] = true

		return
	}

	for i := 0; i < len(nums); i++ {
		if v, ok := (*numsMap)[nums[i]]; ok {
			if v <= 0 {
				continue
			}
		}

		c := make([]int, len(*m))
		copy(c, *m)
		*m = append(*m, nums[i])
		(*numsMap)[nums[i]] -= 1
		traversal(nums, m, numsMap, foundMap, charMap, ret)
		*m = make([]int, len(c))
		copy(*m, c)
		(*numsMap)[nums[i]] += 1
	}

}

//TODO
func permuteUnique(nums []int) [][]int {
	var m []int
	numMap := make(map[int]int)
	foundMap := make(map[string]bool)
	charMap := make(map[int]string)
	var ret [][]int
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] += 1
		charMap[nums[i]] = strconv.Itoa(nums[i])
	}

	traversal(nums, &m, &numMap, &foundMap, &charMap, &ret)
	return ret
}
