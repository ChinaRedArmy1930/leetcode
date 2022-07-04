package main

import "fmt"

func checkOk(queen [][]string, row, col int) bool {
	//检查同一行
	for i := 0; i < len(queen); i++ {
		if queen[row][i] == "Q" {
			return false
		}
	}

	//检查同一列
	for i := 0; i < len(queen); i++ {
		if queen[i][col] == "Q" {
			return false
		}
	}

	//检查左上
	i := row - 1
	j := col - 1
	for i >= 0 && j >= 0 {
		if queen[i][j] == "Q" {
			return false
		}

		i--
		j--
	}

	//检查右下
	i = row + 1
	j = col + 1
	for i < len(queen) && j < len(queen) {
		if queen[i][j] == "Q" {
			return false
		}
		i++
		j++
	}

	//检查右上
	i = row - 1
	j = col + 1
	for i >= 0 && j < len(queen) {
		if queen[i][j] == "Q" {
			return false
		}
		i--
		j++
	}
	//检查左下
	i = row + 1
	j = col - 1
	for i < len(queen) && j >= 0 {
		if queen[i][j] == "Q" {
			return false
		}
		i++
		j--
	}

	return true

}

func traversal(row int, queen *[][]string, res *[][]string) {
	//满足条件, 将queen添加到res中
	if row >= len(*queen) {
		var one []string
		for i := 0; i < len(*queen); i++ {
			fmt.Println((*queen)[i])
			var t string
			for j := 0; j < len((*queen)[i]); j++ {
				t += (*queen)[i][j]
			}
			one = append(one, t)
		}
		*res = append(*res, one)
		fmt.Println()
		return
	}

	for i := 0; i < len((*queen)[row]); i++ {
		if !checkOk(*queen, row, i) {
			continue
		}
		(*queen)[row][i] = "Q"
		traversal(row+1, queen, res)
		(*queen)[row][i] = "."
	}
}

func solveNQueens(n int) [][]string {
	var ans [][]string
	var queen [][]string

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
	fmt.Println(solveNQueens(4))
}
