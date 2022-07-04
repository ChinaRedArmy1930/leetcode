package main

import "fmt"

func checkOk(queen [][]string, row, col int) bool {
	//左上 左下 右上 右下 同一行 同一列
	i := 0
	j := 0

	//左上
	i = row - 1
	j = col - 1
	for i >= 0 && j >= 0 {
		if queen[i][j] == "Q" {
			return false
		}

		i--
		j--
	}

	//左下
	i = row + 1
	j = col - 1
	for i < len(queen) && j >= 0 {
		if queen[i][j] == "Q" {
			return false
		}
		i++
		j--
	}

	//右上
	i = row - 1
	j = col + 1
	for i >= 0 && j < len(queen) {
		if queen[i][j] == "Q" {
			return false
		}
		i--
		j++
	}

	//右下
	i = row + 1
	j = col + 1
	for i < len(queen) && j < len(queen) {
		if queen[i][j] == "Q" {
			return false
		}
		i++
		j++
	}

	//同一行
	for i := 0; i < len(queen[row]); i++ {
		if queen[row][i] == "Q" {
			return false
		}
	}

	//同一列
	for i := 0; i < len(queen); i++ {
		if queen[i][col] == "Q" {
			return false
		}
	}
	return true
}

func traversal(row int, queen *[][]string, ans *int) {
	//满足条件, 添加结果
	if row == len(*queen) {
		*ans += 1
		for i := 0; i < len(*queen); i++ {
			fmt.Println((*queen)[i])
		}
		fmt.Println()
		return
	}

	for i := 0; i < len(*queen); i++ {
		if !checkOk(*queen, row, i) {
			continue
		}
		(*queen)[row][i] = "Q"
		traversal(row+1, queen, ans)
		(*queen)[row][i] = "."
	}
}

func totalNQueens(n int) int {
	var queen [][]string
	ans := 0

	for i := 0; i < n; i++ {
		var t []string
		for j := 0; j < n; j++ {
			t = append(t, ".")
		}
		queen = append(queen, t)
	}

	traversal(0, &queen, &ans)
	return ans
}

func main() {
	fmt.Println(totalNQueens(4))
}
