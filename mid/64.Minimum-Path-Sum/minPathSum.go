package main

import (
	"fmt"
	"sort"
)

func min(x ...int) int {
	sort.Ints(x)
	return x[0]
}

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))

	for i := 0; i < len(dp); i++ {
		t := make([]int, len(grid[i]))
		dp[i] = t
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i == 0 {
				if j == 0 {
					dp[i][j] = grid[i][j]
				} else {
					dp[i][j] += grid[i][j] + dp[i][j-1]
				}
			}

			if j == 0 {
				if i == 0 {
					dp[i][j] = grid[i][j]
				} else {
					dp[i][j] += grid[i][j] + dp[i-1][j]
				}
			}
		}
	}
	/*
		for _, ints := range grid {
			fmt.Println(ints)
		}
		fmt.Println()

		for _, ints := range dp {
			fmt.Println(ints)
		}
		fmt.Println()
	*/
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
