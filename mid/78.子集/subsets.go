package main

import "fmt"

func traversal(index int, arr *[]int, ret *[][]int, num []int) {
	tmp := make([]int, len(*arr))
	copy(tmp, *arr)
	*ret = append(*ret, tmp)

	if index >= len(num) {
		return
	}
	for i := index; i < len(num); i++ {
		*arr = append(*arr, num[i])
		traversal(i+1, arr, ret, num)
		*arr = (*arr)[:len(*arr)-1]
	}
}

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	arr := make([]int, 0)
	traversal(0, &arr, &result, nums)
	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
