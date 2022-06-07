package main

func checkInclusion(s1 string, s2 string) bool {
	left := 0
	right := 0
	valid := 0
	windows := make(map[int32]int, 0)
	need := make(map[int32]int, 0)

	for _, v := range s1 {
		need[v] += 1
	}

	for right < len(s2) {
		c := s2[right]
		right++

		if _, ok := need[int32(c)]; ok {
			windows[int32(c)] += 1
			if windows[int32(c)] == need[int32(c)] {
				valid++
			}
		}

		for valid == len(need) {
			d := s2[left]

			if right-left == len(s1) {
				return true
			}
			left++
			if _, ok := need[int32(d)]; ok {
				if windows[int32(d)] == need[int32(d)] {
					valid--
				}
				windows[int32(d)] -= 1
			}
		}
	}

	return false
}
