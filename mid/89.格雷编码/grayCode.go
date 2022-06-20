package main

import (
	"fmt"
	"math"
	"time"
)

func count(n int) int {
	c := 0
	for i := 0; i < 32; i++ {
		if n&0x1 == 1 {
			c++
		}
		n = n >> 1
	}

	return c
}

func checkOk(nums []int) bool {
	if len(nums) < 1 {
		return false
	}

	fmt.Println(nums)

	t1 := count(nums[0]) ^ count(nums[len(nums)-1])
	if t1 != 1 {
		return false
	}

	return true
}

func traversal(index int, arr *[]int, ret *[]int, leftNum map[int]bool, max int, found *bool) {
	if *found {
		return
	}
	if index >= max {
		//*ret = make([]int, len(*arr))
		//copy(*ret, *arr)
		//*found = true
		if checkOk(*arr) {
			*ret = make([]int, len(*arr))
			copy(*ret, *arr)
			*found = true
		}
		return
	}

	for i, v := range leftNum {
		if !v {
			continue
		}

		if count(((*arr)[len(*arr)-1])^i) != 1 {
			continue
		}

		*arr = append(*arr, i)

		leftNum[i] = false
		traversal(index+1, arr, ret, leftNum, max, found)
		*arr = (*arr)[:len(*arr)-1]
		leftNum[i] = true
	}
}

//todo
func grayCode(n int) []int {
	ret := make([]int, 0)
	arr := make([]int, 0)
	arr = append(arr, 0)
	leftNum := make(map[int]bool, 0)
	max := int(math.Pow(2, float64(n)))
	for i := 1; i < max; i++ {
		leftNum[i] = true
	}
	found := false
	traversal(1, &arr, &ret, leftNum, max, &found)
	return ret
}

func main() {
	//for i := 0; i < 16; i++ {
	//	fmt.Println(grayCode(i))
	//}
	now := time.Now()
	fmt.Println(grayCode(7))
	fmt.Println(time.Now().Sub(now).Seconds())
	//for i := 0; i < 32; i++ {
	//	fmt.Println(count(i))
	//}

}
