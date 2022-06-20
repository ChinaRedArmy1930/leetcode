package main

import (
	"fmt"
	"sort"
)

func max(x ...int) int {
	sort.Ints(x)
	return x[len(x)-1]
}

func min(x ...int) int {
	sort.Ints(x)
	return x[0]
}

func dp(str1, str2 string, x, y int, dpTable *[][]int, visited *[][]bool) int {
	if x >= len(str1) || y >= len(str2) || x < 0 || y < 0 {
		return 0
	}

	if (*visited)[x][y] {
		return (*dpTable)[x][y]
	}

	if str1[x] == str2[y] {
		(*dpTable)[x][y] = dp(str1, str2, x-1, y-1, dpTable, visited) + 1
	} else {
		//else说明x y对应的字符不组成lcs
		(*dpTable)[x][y] = max(dp(str1, str2, x-1, y, dpTable, visited), dp(str1, str2, x, y-1, dpTable, visited))
	}

	(*visited)[x][y] = true
	return (*dpTable)[x][y]
}

func longestCommonSubsequence(text1 string, text2 string) int {
	//对于s1[0...i-1] 和 s2[0...i-1]他们的LCS长度是dpTable[i]
	dpTable := make([][]int, 0)
	visited := make([][]bool, 0)
	for i := 0; i < len(text1)+1; i++ {
		tmp := make([]int, 0)
		v := make([]bool, 0)
		for j := 0; j < len(text2)+1; j++ {
			tmp = append(tmp, 0)
			v = append(v, false)
		}
		dpTable = append(dpTable, tmp)
		visited = append(visited, v)
	}

	return dp(text1, text2, len(text1)-1, len(text2)-1, &dpTable, &visited)

}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
