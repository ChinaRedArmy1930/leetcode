package main

import (
	"fmt"
	"math"
	"sort"
)

type enevlopeType [][]int

func (e enevlopeType) Len() int {
	return len(e)
}

func (e enevlopeType) Less(i, j int) bool {
	if e[i][0] < e[j][0] {
		return true
	}

	if e[i][0] == e[j][0] {
		if e[i][1] < e[j][1] {
			return false
		} else {
			return true
		}
	}

	return false
}

func (e enevlopeType) Swap(i int, j int) {
	e[i], e[j] = e[j], e[i]
}

func dp(envelopes enevlopeType) int {
	if len(envelopes) == 0 {
		return 0
	}

	if len(envelopes) == 1 {
		return 1
	}

	dpTable := make([]int, 0)
	dpTable = append(dpTable, 1)

	sort.Sort(envelopes)
	max := ^math.MaxInt64
	for i := 1; i < len(envelopes); i++ {
		tmp := ^math.MaxInt64
		tmpeve := make([]int, 2)
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				if tmp < dpTable[j]+1 {
					tmp = dpTable[j] + 1
					copy(tmpeve, envelopes[j])
				}
			}
		}

		if tmp == ^math.MaxInt64 {
			tmp = 1
		}

		if max < tmp {
			max = tmp
		}

		dpTable = append(dpTable, tmp)

	}

	if max == ^math.MaxInt64 {
		max = 1
	}

	return max
}

func maxEnvelopes(envelopes [][]int) int {
	return dp(envelopes)
}

func main() {
	//fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
	//fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}, {1, 1}}))
	fmt.Println(maxEnvelopes([][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}}))
}
