package main

import (
	"fmt"
	"strconv"
)

func Sort(ss []string) string {
	comp := func(str1, str2 string) bool {
		return str1+str2 > str2+str1
	}

	result := ""
	for i := 0; i < len(ss); i++ {
		for k := i + 1; k < len(ss); k++ {
			if comp(ss[i], ss[k]) {
				ss[i], ss[k] = ss[k], ss[i]
			}
		}
	}

	for _, v := range ss {
		result += v
	}

	return result
}

func minNumber(nums []int) string {
	NumsString := make([]string, 0)
	for _, v := range nums {
		NumsString = append(NumsString, strconv.Itoa(v))
	}

	return Sort(NumsString)
}

func main() {
	fmt.Println(minNumber([]int{3, 30, 34, 5, 9}))
}
