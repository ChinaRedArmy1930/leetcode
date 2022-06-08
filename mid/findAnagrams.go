package main

func findAnagrams(s string, p string) []int {
	left := 0
	right := 0
	valid := 0

	result := make([]int, 0)

	window := make(map[int32]int, 0)
	need := make(map[int32]int, 0)

	for _, v := range p {
		need[v] += 1
	}

	for right < len(s) {
		c := s[right]
		right++

		if _, ok := need[int32(c)]; ok {
			window[int32(c)] += 1
			if window[int32(c)] == need[int32(c)] {
				valid++
			}
		}

		for valid == len(need) {
			d := s[left]

			if right-left == len(p) {
				result = append(result, left)
			}

			left++

			if _, ok := need[int32(d)]; ok {
				if window[int32(d)] == need[int32(d)] {
					valid--
				}
				window[int32(d)] -= 1
			}
		}
	}
	return result
}
