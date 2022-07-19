package main

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		t := make([]int, n)
		matrix[i] = t
	}

	action := 1 // 1右  2下 3左 0上
	x := 0
	y := 0
	shang := 0
	xia := n - 1
	zuo := 0
	you := n - 1
	for i := 1; i <= n*n; i++ {
		matrix[x][y] = i
		switch action % 4 {
		case 1:
			y++
			if y == you {
				shang++
				action++
			}
		case 2:
			x++
			if x == xia {
				you--
				action++
			}
		case 3:
			y--
			if y == zuo {
				xia--
				action++
			}
		case 0:
			x--
			if x == shang {
				zuo++
				action++
			}
		}
	}

	return matrix
}

func main() {
	m := generateMatrix(5)
	for _, l := range m {
		fmt.Println(l)
	}
}
