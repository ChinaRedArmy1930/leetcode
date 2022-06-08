package main

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	ans := 0
	window := make(map[int32]int, 0)
	for right < len(s) {
		c := s[right]
		right++
		window[int32(c)] += 1

		for window[int32(c)] > 1 {
			d := s[left]
			left++
			window[int32(d)] -= 1
		}

		if ans < right-left {
			ans = right - left
		}
	}

	return ans
}
