package main

import "fmt"

func main() {
	fmt.Print(findNumberIn2DArray([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	//fmt.Print(findNumberIn2DArray([][]int{{5}}, 5))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	startIndexWight := len(matrix[0]) - 1
	startIndexHeight := 0

	fmt.Printf("%d  %d  %d\n", startIndexWight, startIndexHeight, len(matrix)-1)

	i := startIndexWight
	j := startIndexHeight
	for i >= 0 && j <= len(matrix)-1 {
		fmt.Printf("i = %d  j = %d \n", i, j)
		if target == matrix[j][i] {
			return true
		}

		if target > matrix[j][i] {
			j++
			continue
		}

		if target < matrix[j][i] {
			i--
			continue
		}
	}

	return false
}
