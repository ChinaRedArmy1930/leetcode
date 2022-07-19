package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		t := make([]int, len(obstacleGrid[i]))
		dp[i] = t
	}

	lz := false
	hz := false

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i == 0 {
				if obstacleGrid[i][j] == 1 {
					lz = true
				}

				if !lz {
					dp[i][j] = 1
				}

			}

			if j == 0 {
				if obstacleGrid[i][j] == 1 {
					hz = true
				}

				if !hz {
					dp[i][j] = 1
				}

			}

			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			}
		}
	}

	for _, i := range dp {
		fmt.Println(i)
	}

	fmt.Println()

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	for _, i := range dp {
		fmt.Println(i)
	}

	fmt.Println()
	return dp[len(dp)-1][len(dp[0])-1]
}

func main() {
	//fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	//fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
	//fmt.Println(uniquePathsWithObstacles([][]int{{1, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{1}, {0}}))
}
