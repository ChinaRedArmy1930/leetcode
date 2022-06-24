package main

import (
	"fmt"
	"sort"
)

func min(x ...int) int {
	sort.Ints(x)
	return x[0]
}

func max(x ...int) int {
	sort.Ints(x)
	return x[len(x)-1]
}

//TODO 自顶向上
func minDistance2(word1, word2 string) int {
	dpTable := make([][]int, len(word1)+1)
	//dpMap := make([][]bool, len(word1)+1)

	tmp := make([]int, 0)
	for j := 0; j < len(word2)+1; j++ {
		tmp = append(tmp, j)
	}

	dpTable[0] = tmp
	tmp = make([]int, len(word1)+1)
	for i := 1; i < len(word1)+1; i++ {
		tmp = make([]int, len(word2)+1)
		tmp[0] = i
		dpTable[i] = tmp
	}

	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dpTable[i+1][j+1] = dpTable[i][j]
			} else {
				dpTable[i+1][j+1] = min(dpTable[i+1][j], dpTable[i][j], dpTable[i][j+1]) + 1
			}
		}
	}

	for i := 0; i < len(dpTable); i++ {
		fmt.Println(dpTable[i])
	}
	return dpTable[len(dpTable)-1][len(dpTable[0])-1]
}

func dp(dpTable [][]int, x, y int, word1, word2 string, dpMap *[][]bool) int {
	if x < 0 {
		return y + 1
	}

	if y < 0 {
		return x + 1
	}

	if (*dpMap)[x][y] {
		return dpTable[x][y]
	}

	if word1[x] == word2[y] {
		return dp(dpTable, x-1, y-1, word1, word2, dpMap)
	} else {
		in := dp(dpTable, x-1, y, word1, word2, dpMap) + 1
		del := dp(dpTable, x, y-1, word1, word2, dpMap) + 1
		rep := dp(dpTable, x-1, y-1, word1, word2, dpMap) + 1
		dpTable[x][y] = min(in, del, rep)
	}

	(*dpMap)[x][y] = true
	return dpTable[x][y]

}

//自顶向下
func minDistance(word1 string, word2 string) int {
	dpTable := make([][]int, len(word1))
	dpMap := make([][]bool, len(word1))
	for i := 0; i < len(word1); i++ {
		t := make([]int, len(word2))
		t2 := make([]bool, len(word2))
		for j := 0; j < len(word2); j++ {
			t[j] = 0
			t2[j] = false
		}
		dpTable[i] = t
		dpMap[i] = t2
	}

	return dp(dpTable, len(word1)-1, len(word2)-1, word1, word2, &dpMap)
}

func main() {
	fmt.Println(minDistance2("rad", "apple"))
	fmt.Println(minDistance("rad", "apple"))
}
