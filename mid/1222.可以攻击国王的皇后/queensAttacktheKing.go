package main

import "fmt"

func check(queen []int, king []int, checked *[]bool) bool {
	kingRow := king[0]
	kingCol := king[1]
	queenRow := queen[0]
	queenCol := queen[1]

	i1 := kingRow + 1
	j1 := kingCol + 1
	i2 := queenRow
	j2 := queenCol

	//右下
	if i1 < 8 && j1 < 8 && (*checked)[0] == false {
		if i1 == i2 && j1 == j2 {
			(*checked)[0] = true
			return true
		}
		i1++
		j1++
	}

	//左上
	i1 = kingRow - 1
	j1 = kingCol - 1
	if i1 >= 0 && j1 >= 0 && (*checked)[1] == false {
		if i1 == i2 && j1 == j2 {
			(*checked)[1] = true
			return true
		}
		i1--
		j1--
	}

	//右上
	i1 = kingRow - 1
	j1 = kingCol + 1
	if i1 >= 0 && j1 < 8 && (*checked)[2] == false {
		if i1 == i2 && j1 == j2 {
			(*checked)[2] = true
			return true
		}
		i1--
		j1++
	}

	//左下
	i1 = kingRow + 1
	j1 = kingCol - 1
	if i1 < 8 && j1 >= 0 && (*checked)[3] == false {
		if i1 == i2 && j1 == j2 {
			(*checked)[3] = true
			return true
		}
		i1++
		j1--
	}

	//上
	i1 = kingRow - 1
	j1 = kingCol
	for i1 >= 0 {
		if i1 == i2 && j1 == j2 && (*checked)[4] == false {
			(*checked)[4] = true
			return true
		}
		i1--
	}

	//下
	i1 = kingRow + 1
	j1 = kingCol
	for i1 < 8 {
		if i1 == i2 && j1 == j2 && (*checked)[5] == false {
			(*checked)[5] = true
			return true
		}
		i1++
	}

	//左
	i1 = kingRow
	j1 = kingCol - 1
	for j1 >= 0 {
		if i1 == i2 && j1 == j2 {
			return true
		}
		j1--
	}

	//右
	i1 = kingRow
	j1 = kingCol + 1
	for j1 < 8 {
		if i1 == i2 && j1 == j2 {
			return true
		}
		j1++
	}

	return false
}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	var ans [][]int
	checked := make([]bool, 8)
	for i := 0; i < len(checked); i++ {
		checked[i] = false
	}
	for _, queen := range queens {
		if check(king, queen, &checked) {
			ans = append(ans, queen)
		}
	}

	return ans
}

func main() {
	fmt.Println(queensAttacktheKing([][]int{[]int{0, 1}, []int{1, 0}, []int{4, 0}, []int{0, 4}, []int{3, 3}, []int{2, 4}}, []int{0, 0}))
}
