package main

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	window := make(map[string]int)
	max := 1
	left := 0
	right := 0
	valid := 0
	for right < len(s) {
		c := s[right]
		if _, ok := window[string(c)]; !ok {
			valid++
		}
		window[string(c)] += 1
		right++

		for valid > 2 && left < len(s) {
			d := s[left]
			left++

			window[string(d)] -= 1

			if window[string(d)] == 0 {
				delete(window, string(d))
			}

			if _, ok := window[string(d)]; !ok {
				valid--
			}
		}

		if right-left > max {
			fmt.Println(s[left:right])
			max = right - left
		}
	}

	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("aa"))
}
