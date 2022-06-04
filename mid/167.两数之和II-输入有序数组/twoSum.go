package main

func twoSum2(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		}

		if numbers[l]+numbers[r] > target {
			r = r - 1
		} else {
			l = l + 1
		}
	}

	return []int{-1, -1}
}

func twoSum(numbers []int, target int) []int {
	l := 0
	r := 0
	for r < len(numbers) {
		l = r + 1
		for l < len(numbers) {
			if numbers[l]+numbers[r] == target {
				if l < r {
					return []int{l + 1, r + 1}
				} else {
					return []int{r + 1, l + 1}
				}

			}
			l++
		}
		r++
	}

	return []int{-1, -1}
}
