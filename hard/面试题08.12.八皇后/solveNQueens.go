package main

import "fmt"

func QueenOk(queen *[][]string, row, col int) bool {
	if row >= len(*queen) {
		//return false
	}

	if col >= len((*queen)[0]) {
		//return false
	}

	//检查同一列
	for i := 0; i < len(*queen); i++ {
		if (*queen)[i][col] == "Q" {
			return false
		}
	}

	//同一行
	for j := 0; j < len((*queen)[col]); j++ {
		if (*queen)[row][j] == "Q" {
			return false
		}
	}

	//同一对角线 左上
	i := row - 1
	j := col - 1
	for i >= 0 && j >= 0 {
		if (*queen)[i][j] == "Q" {
			return false
		}
		i--
		j--
	}

	//右下
	i = row + 1
	j = col + 1
	for i < len(*queen) && j < len((*queen)[0]) {
		if (*queen)[i][j] == "Q" {
			return false
		}
		i++
		j++
	}

	//右上
	i = row - 1
	j = col + 1
	for i >= 0 && j < len((*queen)[0]) {
		if (*queen)[i][j] == "Q" {
			return false
		}
		i--
		j++
	}

	//左下
	i = row + 1
	j = col - 1
	for i < len(*queen) && j >= 0 {
		if (*queen)[i][j] == "Q" {
			return false
		}
		i++
		j--
	}
	return true
}

func traversal(row int, ret *[][]string, queen *[][]string) {
	//满足该条件就return
	if row >= len(*queen) {
		var ans []string
		for i := 0; i < len(*queen); i++ {
			fmt.Println((*queen)[i])
			line := ""
			for j := 0; j < len((*queen)[i]); j++ {
				line += (*queen)[i][j]
			}
			ans = append(ans, line)
		}
		fmt.Println()
		*ret = append(*ret, ans)

		return
	}

	//执行某项操作
	for i := 0; i < len((*queen)[row]); i++ {
		if !QueenOk(queen, row, i) {
			continue
		}
		(*queen)[row][i] = "Q"
		traversal(row+1, ret, queen)
		(*queen)[row][i] = "."
	}
}

func solveNQueens(n int) [][]string {
	//初始化数组
	var queens [][]string
	var ans [][]string
	for i := 0; i < n; i++ {
		var t []string
		for j := 0; j < n; j++ {
			t = append(t, ".")
		}
		queens = append(queens, t)
	}

	traversal(0, &ans, &queens)

	return ans
}

func main() {
	fmt.Println(solveNQueens(4))
}
