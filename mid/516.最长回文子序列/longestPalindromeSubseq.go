package main

import (
	"fmt"
	"sort"
)

func max(x ...int) int {
	sort.Ints(x)
	return x[len(x)-1]
}

func longestPalindromeSubseq(s string) int {
	dpTable := make([][]int, len(s))

	for i := 0; i < len(s); i++ {
		t := make([]int, len(s))
		for j := 0; j < len(s); j++ {
			if i == j {
				t[j] = 1
			}

			if i > j {
				t[j] = 0
			}

			if i < j {
				t[j] = -1
			}
		}
		dpTable[i] = t
	}

	//斜着遍历
	for k := 0; k < len(s); k++ {
		i := 0
		j := k + 1
		for i < len(s) && j < len(s) {
			if s[i] == s[j] {
				dpTable[i][j] = dpTable[i+1][j-1] + 2
			} else {
				dpTable[i][j] = max(dpTable[i][j-1], dpTable[i+1][j])
			}
			i++
			j++
		}
	}

	for i := 0; i < len(dpTable); i++ {
		fmt.Println(dpTable[i])
	}

	return dpTable[0][len(s)-1]
}

func main() {
	longestPalindromeSubseq("bbbab")
}
