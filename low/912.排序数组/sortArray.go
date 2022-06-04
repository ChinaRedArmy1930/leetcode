package main

func SelectionArray(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		value := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if value > nums[j] {
				value = nums[j]
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

func Merge(nums []int, start, mid, end int) {
	A := nums[start : mid+1]
	B := nums[mid+1 : end+1]

	i := 0
	j := 0

	tmp := make([]int, 0)

	if len(A) == 0 {
		copy(nums, B)
		return
	} else if len(B) == 0 {
		copy(nums, A)
		return
	}

	for {
		if i >= len(A) || j >= len(B) {
			if i > mid {
				tmp = append(tmp, B[j:]...)
				break
			} else {
				tmp = append(tmp, A[i:]...)
				break
			}
		}

		if A[i] < B[j] {
			tmp = append(tmp, A[i])
			i++
			continue
		}

		if A[i] > B[j] {
			tmp = append(tmp, B[j])
			j++
			continue
		}

		if A[i] == B[j] {
			tmp = append(tmp, []int{A[i], B[j]}...)
			i++
			j++
		}
	}

	copy(nums[start:end+1], tmp)
}

func MergeSortC(nums []int, start, end int) {
	if start >= end-1 {
		if nums[start] > nums[end] {
			nums[start], nums[end] = nums[end], nums[start]
		}

		return
	}

	mid := (start + end) / 2
	MergeSortC(nums, start, mid)
	MergeSortC(nums, mid+1, end)
	Merge(nums, start, mid, end)

}

func MergeSort(nums []int) []int {
	MergeSortC(nums, 0, len(nums)-1)
	return nums
}

func InsertSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		value := nums[i]
		j := i - 1
		//倒序非常重要
		for j = i - 1; j >= 0; j-- {
			if nums[j] > value {
				nums[j+1] = nums[j]
			} else {
				//非常重要, 得到分界点
				break
			}
		}
		nums[j+1] = value
	}

	return nums
}
