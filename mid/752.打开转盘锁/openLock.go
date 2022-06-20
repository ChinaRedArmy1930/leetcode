package main

import "container/list"

func addOne(str string, index int) string {
	strArr := make([]uint8, 0)
	for k, v := range str {
		if k == index {
			if v == '9' {
				strArr = append(strArr, '0')
			} else {
				strArr = append(strArr, uint8(v+1))
			}
		} else {
			strArr = append(strArr, uint8(v))
		}
	}
	ret := ""
	for _, v := range strArr {
		ret += string(v)
	}

	return ret
}

func subOne(str string, index int) string {
	strArr := make([]uint8, 0)
	for k, v := range str {
		if k == index {
			if v == '0' {
				strArr = append(strArr, '9')
			} else {
				strArr = append(strArr, uint8(v-1))
			}
		} else {
			strArr = append(strArr, uint8(v))
		}
	}
	ret := ""
	for _, v := range strArr {
		ret += string(v)
	}

	return ret
}

func openLock(deadends []string, target string) int {
	str := list.New()
	str.PushBack("0000")
	visited := make(map[string]bool)
	visited["0000"] = true
	count := 0

	for str.Len() > 0 {
		l := str.Len()
		for i := 0; i < l; i++ {
			t := str.Front().Value.(string)
			str.Remove(str.Front())
			if t == target {
				return count
			}

			foundDead := false
			for _, deadend := range deadends {
				if t == deadend {
					foundDead = true
					break
				}
			}

			if foundDead {
				continue
			}

			for j := 0; j < 4; j++ {
				add := addOne(t, j)
				if _, ok := visited[add]; !ok {
					visited[addOne(t, j)] = true
					str.PushBack(addOne(t, j))
				}

				sub := subOne(t, j)
				if _, ok := visited[sub]; !ok {
					visited[subOne(t, j)] = true
					str.PushBack(subOne(t, j))
				}
			}
		}
		count++
	}

	return -1
}

func main() {

}
