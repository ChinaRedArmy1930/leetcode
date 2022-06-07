package main

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	result := 0
	window := make(map[int32]int, 0)
	for right < len(s) {
		c := s[right]
		right++
		window[int32(c)] += 1

		v, _ := window[int32(c)]
		for v > 1 {
			d := s[left]
			left++
			window[int32(d)] -= 1
			v, _ = window[int32(c)]
		}

		if right-left > result {
			result = right - left
		}

	}

	return result
}
