package main

import (
	"fmt"
)

func dp(dpTable [][]int, i, j int, s, t string, visited [][]bool) int {
	if i >= len(s) || j >= len(t) {
		return dpTable[i][j]
	}

	if visited[i][j] == true {
		return dpTable[i][j]
	}

	if s[i] == t[j] {
		dpTable[i][j] = dp(dpTable, i+1, j+1, s, t, visited) + dp(dpTable, i+1, j, s, t, visited)
	} else {
		dpTable[i][j] = dp(dpTable, i+1, j, s, t, visited)
	}
	visited[i][j] = true
	return dpTable[i][j]
}

func numDistinct(s string, t string) int {
	//dp[i][j] 代表s[0...i] t[0...j] s的子序列在t中出现的个数
	//base case :当 i < 0时, 子序列的个数为0
	dpTable := make([][]int, len(s)+1)
	visited := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		tmp := make([]int, len(t)+1)
		t2 := make([]bool, len(t)+1)
		for j := 0; j < len(t)+1; j++ {
			tmp[j] = 0
			t2[j] = false
		}
		dpTable[i] = tmp
		visited[i] = t2
	}

	for i := 0; i < len(s)+1; i++ {
		dpTable[i][len(t)] = 1
		visited[i][len(t)] = true
	}

	for i := 0; i < len(t)+1; i++ {
		visited[len(s)][i] = true
	}

	//for i := 0; i < len(dpTable); i++ {
	//	fmt.Println(dpTable[i])
	//}
	//fmt.Println()
	dp(dpTable, 0, 0, s, t, visited)

	//for i := 0; i < len(dpTable); i++ {
	//	fmt.Println(dpTable[i])
	//}
	return dpTable[0][0]
}

func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println(numDistinct("babgbag", "bag"))
}
