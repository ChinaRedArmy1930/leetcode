package main

import (
	"fmt"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	if len(score) == 0 {
		return nil
	}

	indexMap := make(map[int]int, len(score))
	result := make([]string, len(score))
	for k, v := range score {
		indexMap[v] = k
	}

	for i := 1; i < len(score); i++ {
		value := score[i]
		j := i - 1
		for j = i - 1; j >= 0; j-- {
			if value >= score[j] {
				score[j+1] = score[j]
			} else {
				break
			}
		}
		score[j+1] = value
	}

	for k, v := range score {
		switch k {
		case 0:
			result[indexMap[v]] = "Gold Medal"
		case 1:
			result[indexMap[v]] = "Silver Medal"
		case 2:
			result[indexMap[v]] = "Bronze Medal"
		default:
			result[indexMap[v]] = strconv.Itoa(k + 1)
		}

	}

	return result
}

func main() {
	testCase := [][]int{
		{10, 3, 8, 9, 4},
		//{7, 1, 3, 2, 4, 5, 6, 12},
	}
	for _, v := range testCase {
		fmt.Println(findRelativeRanks(v))
	}

}
