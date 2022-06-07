package main

import "math"

func needShrink(w map[int32]int) bool {
	OnePCount := 0
	for _, v := range w {
		if v == 0 {

			return false
		}
		if v > 1 {
			OnePCount += 1
		}
	}

	return OnePCount >= 1
}

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
	lens := math.MaxInt32

	windows := make(map[int32]int, 0)
	need := make(map[int32]int, 0)

	for _, v := range t {
		need[v] += 1 //注意: 不是赋值为1 而是+=1
	}

	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[int32(c)]; ok {
			windows[int32(c)] += 1
			if windows[int32(c)] == need[int32(c)] {
				valid++
			}
		}

		for valid == len(need) { //注意, 是need的大小  不是t的大小
			d := s[left]

			if lens > right-left {
				start = left
				lens = right - left
			}

			left++

			if _, ok := need[int32(d)]; ok {
				if windows[int32(d)] == need[int32(d)] {
					valid--
				}
				windows[int32(d)] -= 1 //注意 等对比完了再减1
			}
		}
	}

	if lens == math.MaxInt32 {
		return ""
	}

	return s[start : start+lens]
}
