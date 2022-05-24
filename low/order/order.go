package main

import "fmt"

func BubbleOrder(arrays []int) []int {
	arrayLen := len(arrays)
	for i := 0; i < arrayLen; i++ {
		for j := i; j < arrayLen; j++ {
			if arrays[i] < arrays[j] {
				arrays[i], arrays[j] = arrays[j], arrays[i]
			}
		}
	}

	return arrays
}

func InsertOrder(nums []int) []int {
	numsLen := len(nums)

	if numsLen == 0 {
		return []int{}
	}

	for i := 0; i < numsLen; i++ {
		value := nums[i]
		j := 0
		for j = i - 1; j >= 0; j-- {
			if value >= nums[j] {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}

		nums[j+1] = value
	}

	return nums
}

func SelectionOrder(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		index := i + 1

		for j := i + 1; j < len(nums); j++ {
			if nums[index] > nums[j] {
				index = j
			}
		}

		if index != i+1 {
			nums[i+1], nums[index] = nums[index], nums[i+1]
		}
	}
	return nums
}

func MergeNums(array []int, start, mid, end int) {
	fmt.Printf("======== stat = %d  mid = %d  end = %d \n", start, mid, end)
	A := array[start:mid]
	B := array[mid:end]
	fmt.Println(A, B)
	result := make([]int, 0)
	i, j := 0, 0

	if len(A) == 0 {
		copy(array, B)
		return
	} else if len(B) == 0 {
		copy(array, A)
		return
	}

	for {
		if i >= len(A) || j >= len(B) {
			if i > j {
				result = append(result, B[j:]...)
			} else {
				result = append(result, A[i:]...)
			}
			break
		}

		if A[i] < B[j] {
			result = append(result, A[i])
			i++
		} else if A[i] >= B[j] {
			result = append(result, B[j])
			j++
		}
	}

	copy(array, result)
	fmt.Println(result)
}

func MergeSortC(start int, end int, array []int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2

	MergeSortC(start, mid, array)
	MergeSortC(mid+1, end, array)
	MergeNums(array, start, mid, end)
}

func MergeOrder(nums []int) []int {
	MergeSortC(0, len(nums)-1, nums)
	return nums
}

func main() {
	arrays := []int{1, 4, 3, 5, 6, 3, 2, 7, 6, 9, 10, 8}
	//fmt.Println(BubbleOrder(arrays))
	fmt.Println(MergeOrder(arrays))
	//arrays2 := []int{1, 4}
	//MergeNums(arrays2, 0, 0, 1)
	//fmt.Println(arrays2)
}
