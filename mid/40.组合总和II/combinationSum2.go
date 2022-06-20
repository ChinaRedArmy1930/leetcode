package main

import (
	"fmt"
	"time"
)

func checkArrayEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] != b[j] {
			return false
		}
		i++
		j++
	}

	return true

	//a1 := make(map[int]int)
	//b1 := make(map[int]int)
	//
	//if len(a) != len(b) {
	//	return false
	//}
	//
	//for _, v := range a {
	//	a1[v] += 1
	//}
	//
	//for _, v := range b {
	//	b1[v] += 1
	//}
	//
	//for k, v := range a1 {
	//	if b1[k] != v {
	//		return false
	//	}
	//}

	return true
}

//true 代表可以终止此次循环了 false代表继续循环
func traversal(index int, candidates []int, target int, sum *[]int, ret *[][]int) {
	tmp := 0

	for _, v := range *sum {
		tmp += v

	}
	if tmp > target {
		return
	}

	if index >= len(candidates) || tmp == target {

		if tmp == target {
			input := true

			for _, v := range *ret {
				if checkArrayEqual(*sum, v) {
					input = false
				}
			}
			if !input {
				return
			}
			tmp := make([]int, len(*sum))
			copy(tmp, *sum)
			*ret = append(*ret, tmp)

		}

		return
	}

	for i := index; i < len(candidates); i++ {
		if candidates[i] > target {
			return
		}
		//TODO 这里是为啥？？？？？？？？？？？？？？？？？？？？？？？？？？？？？？？？
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}

		*sum = append(*sum, candidates[i])
		traversal(i+1, candidates, target, sum, ret)
		*sum = (*sum)[:len(*sum)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sum := make([]int, 0)
	//先排序
	for i := 1; i < len(candidates); i++ {
		value := candidates[i]
		j := i - 1
		for j = i - 1; j >= 0; j-- {
			if value < candidates[j] {
				candidates[j+1] = candidates[j]
			} else {
				break
			}
		}
		candidates[j+1] = value
	}

	//fmt.Println(candidates)

	traversal(0, candidates, target, &sum, &result)
	return result
}

func main() {
	for _, v := range combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8) {
		fmt.Println(v)
	}

	for _, v := range combinationSum2([]int{2, 5, 2, 1, 2}, 5) {
		fmt.Println(v)
	}

	for _, v := range combinationSum2([]int{1}, 2) {
		fmt.Println(v)
	}

	now := time.Now()

	for _, v := range combinationSum2([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		30) {
		fmt.Println(v)
	}
	//
	fmt.Println(time.Now().Sub(now).Seconds())
	//fmt.Println(checkArrayEqual([]int{1, 2, 5}, []int{1, 1, 6}))
	//fmt.Println(1 ^ 2 ^ 5)
}
