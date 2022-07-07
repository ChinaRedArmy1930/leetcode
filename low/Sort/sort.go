package main

import (
	"fmt"
)

func BubbleSort(arrays []int) []int {
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

func InsertSort(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		value := nums[i]
		for j = i - 1; j >= 0; j-- {
			if value > nums[j] {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = value
	}

	return nums
}

func SelectionSort(nums []int) []int {
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
	A := array[start : mid+1]
	B := array[mid+1 : end+1]
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

	copy(array[start:end+1], result)
	fmt.Println(result)
}

func MergeSortC(start int, end int, array []int) {
	if start >= end-1 {
		if array[start] > array[end] {
			array[start], array[end] = array[end], array[start]
		}

		return
	}

	mid := (start + end) / 2

	MergeSortC(start, mid, array)
	MergeSortC(mid+1, end, array)
	MergeNums(array, start, mid, end)
}

func MergeSort(nums []int) []int {
	MergeSortC(0, len(nums)-1, nums)
	return nums
}

func ModifyHeap(nums []int, start int, heapSize int) {
	left := 2*start + 1
	right := left + 1
	largest := start
	if left < heapSize && nums[left] > nums[largest] {
		largest = left
	}

	if right < heapSize && nums[right] > nums[largest] {
		largest = right
	}

	if start != largest {
		nums[start], nums[largest] = nums[largest], nums[start]
		ModifyHeap(nums, largest, heapSize)
	}
}

func heapSort(nums []int) []int {
	//作用:构建堆中三个
	buildMaxHeap := func(nums []int) {
		for i := (len(nums) / 2) - 1; i >= 0; i-- {
			ModifyHeap(nums, i, len(nums))
		}
	}
	_ = buildMaxHeap
	//构建大顶堆,确保每一个三角形堆顶是三角形的三个数中最大的
	buildMaxHeap(nums)

	//从第一个数开始, 将堆顶的数字放到末尾
	for i := 0; i < len(nums); i++ {
		nums[0], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[0]
		ModifyHeap(nums, 0, len(nums)-i-1)
	}
	return nums
}

func partition(nums []int, start int, end int) int {
	left := start + 1
	right := end
	value := nums[start]

	for left < right {
		for left < right && nums[left] < value {
			left++
		}

		for left < right && nums[right] > value {
			right--
		}

		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	if left == right && nums[right] > value {
		right--
	}

	nums[start], nums[right] = nums[right], nums[start]
	return right
}

func quickSortR(nums []int, start int, end int) {
	if start > end {
		return
	}

	middle := partition(nums, start, end)
	quickSortR(nums, start, middle-1)
	quickSortR(nums, middle+1, end)
}

func quickSort1(nums []int) []int {
	quickSortR(nums, 0, len(nums)-1)
	return nums
}

func main() {
	//arrays := []int{1, 4, 3, 5, 6, 3, 2, 7, 6, 9, 10, 8}
	arrays := []int{8, 6, 4, 5, 7, 3, 9, 10}
	//fmt.Println(BubbleSort(arrays))
	//fmt.Println(MergeSort(arrays))
	//fmt.Println(InsertSort(arrays))
	//fmt.Println(heapSort(arrays))
	//arrays2 := []int{1, 4}
	//MergeNums(arrays2, 0, 0, 1)
	//fmt.Println(arrays2)
	fmt.Println(quickSort1(arrays))
}
