package main

import "fmt"

/*

输入：words = ["bella","label","roller"]
输出：["e","l","l"]

*/

func getCommon(m1, m2 map[string]int) map[string]int {
	ret := make(map[string]int)
	for s, i := range m1 {
		if m2[s] >= i {
			ret[s] = i
		}

		if m2[s] >= 1 && m2[s] < i {
			ret[s] = m2[s]
		}
	}

	return ret
}

func commonChars(words []string) []string {
	found := make(map[string]int)

	for _, v := range words[0] {
		found[string(v)] += 1
	}

	for i := 1; i < len(words); i++ {
		tmp := make(map[string]int)

		for _, v := range words[i] {
			tmp[string(v)] += 1
		}

		found = getCommon(found, tmp)
	}

	ret := make([]string, 0)
	for s, count := range found {
		for i := 0; i < count; i++ {
			ret = append(ret, s)
		}
	}

	return ret
}

func main() {
	//fmt.Println(commonChars([]string{"bella", "label", "roller"}))
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
}
