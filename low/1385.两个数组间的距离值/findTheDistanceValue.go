package main

import "math"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	result := 0
	for i := 0; i < len(arr1); i++ {
		flag := true
		for j := 0; j < len(arr2); j++ {
			if math.Abs(float64(arr1[i]-arr2[j])) <= float64(d) {
				flag = false
				break
			}
		}

		if flag {
			result++
		}
	}

	return result
}
