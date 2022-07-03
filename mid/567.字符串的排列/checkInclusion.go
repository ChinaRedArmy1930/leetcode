package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	window := make(map[string]int)
	need := make(map[string]int)

	for i := 0; i < len(s1); i++ {
		need[string(s1[i])] += 1
	}

	left := 0
	right := 0
	vaild := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := need[string(c)]; ok {
			window[string(c)] += 1
			if window[string(c)] == need[string(c)] {
				vaild++
			}
		}

		for vaild == len(need) {
			if right-left == len(need) {
				return true
			}
			d := s2[left]
			left++
			window[string(c)] += 1
			if _, ok := need[string(d)]; ok {
				if window[string(s2[left])] == need[string(s2[left])] {
					vaild--
				}
				window[string(s2[left])] -= 1
			}
		}
	}

	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}
