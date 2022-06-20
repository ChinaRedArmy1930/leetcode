package main

import (
	"fmt"
	"strconv"
)

func traversal(index int, keys [][]string, tmp *string, ret *[]string) {
	if index >= len(keys) {
		*ret = append(*ret, *tmp)
		return
	}

	for i := 0; i < len(keys[index]); i++ {
		*tmp += keys[index][i]
		traversal(index+1, keys, tmp, ret)
		*tmp = (*tmp)[:len(*tmp)-1]
	}
}

func getKeys(digits string, m map[int][]string) [][]string {
	ret := make([][]string, 0)
	for _, v := range digits {
		num, _ := strconv.Atoi(string(v))
		ret = append(ret, m[num])
	}

	return ret
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	result := make([]string, 0)
	phoneMap := map[int][]string{
		2: []string{"a", "b", "c"},
		3: []string{"d", "e", "f"},
		4: []string{"g", "h", "i"},
		5: []string{"j", "k", "l"},
		6: []string{"m", "n", "o"},
		7: []string{"p", "q", "r", "s"},
		8: []string{"t", "u", "v"},
		9: []string{"w", "x", "y", "z"},
	}
	_ = phoneMap

	keys := getKeys(digits, phoneMap)
	tmp := ""
	traversal(0, keys, &tmp, &result)
	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("7"))
}
