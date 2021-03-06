package main

import (
	"math"
)

/*
  1. 将字符添加到滑动窗口中
	-> 什么样的字符能被添加到滑动窗口?
       在t中的字符 都可以被添加到滑动窗口
  2. 如果滑动窗口满足要求, 则优化该窗口
   ->
*/
func minWindow(s string, t string) string {
	left := 0
	right := 0
	valid := 0

	start := 0
	lens := math.MaxInt64
	window := make(map[int32]int, 0)
	need := make(map[int32]int, 0)

	for _, v := range t {
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

			if lens > right-left {
				start = left
				lens = right - start
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

	if lens == math.MaxInt64 {
		return ""
	}

	return s[start : start+lens]
}
