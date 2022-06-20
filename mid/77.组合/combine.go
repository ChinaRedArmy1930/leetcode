package main

import "fmt"

func traversal(index int, ret *[][]int, arr *[]int, k int, n int) {
	if len(*arr) >= k {
		tmp := make([]int, k)
		copy(tmp, *arr)
		*ret = append(*ret, tmp)
		return
	}

	for i := index; i <= n; i++ {
		*arr = append(*arr, i)
		traversal(i+1, ret, arr, k, n)
		*arr = (*arr)[:len(*arr)-1]
	}
}

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	arr := make([]int, 0)
	traversal(1, &result, &arr, k, n)
	return result
}

func main() {
	fmt.Println(combine(4, 2))
}
