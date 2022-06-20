package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func traversal(matrix [][]int, x, y int, res *[]int, sum *int, min *int) {
	if len(*res) > len(matrix) || x >= len(matrix) || y >= len(matrix) || x < 0 || y < 0 {
		return
	}

	*res = append(*res, matrix[x][y])
	*sum += matrix[x][y]

	if len(*res) == len(matrix) {

		if *min > *sum {
			*min = *sum
		}
	}

	traversal(matrix, x+1, y-1, res, sum, min)
	traversal(matrix, x+1, y, res, sum, min)
	traversal(matrix, x+1, y+1, res, sum, min)

	*res = (*res)[:len(*res)-1]
	*sum -= matrix[x][y]
}

func max(x ...int) int {
	sort.Ints(x)
	return x[len(x)-1]
}

func min(x ...int) int {
	sort.Ints(x)
	return x[0]
}

func dp(matrix [][]int, x, y int, matrixMap *map[string]int) int {
	if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix) {
		return math.MaxInt64
	}

	if x == 0 {
		return matrix[x][y]
	}

	if v, ok := (*matrixMap)[strconv.Itoa(x)+"#"+strconv.Itoa(y)]; ok {
		return v
	}

	a := dp(matrix, x-1, y-1, matrixMap)
	b := dp(matrix, x-1, y, matrixMap)
	c := dp(matrix, x-1, y+1, matrixMap)
	ret := matrix[x][y] + min(a, b, c)

	(*matrixMap)[strconv.Itoa(x)+"#"+strconv.Itoa(y)] = ret

	return ret
}

//动态规划解法
func minFallingPathSumDp(matrix [][]int) int {
	res := math.MaxInt64
	matrixMap := make(map[string]int)
	for i := 0; i < len(matrix); i++ {
		res = min(res, dp(matrix, len(matrix)-1, i, &matrixMap))
	}
	return res
}

//回溯解法
func minFallingPathSumBackTracking(matrix [][]int) int {
	res := make([]int, 0)
	sum := 0
	min := math.MaxInt64
	for i := 0; i < len(matrix); i++ {
		traversal(matrix, 0, i, &res, &sum, &min)
	}

	return min
}

func main() {
	fmt.Println(min([]int{1, 2, 3, 4, 56}...))
}
