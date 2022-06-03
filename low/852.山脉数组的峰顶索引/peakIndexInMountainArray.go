package main

func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr) - 1
	mid := 0
	for start <= end {
		mid = start + (end-start)/2
		if mid-1 >= 0 && mid+1 < len(arr) {
			if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
				return mid
			} else if arr[mid] > arr[mid-1] && arr[mid] < arr[mid+1] {
				start = mid + 1
			} else if arr[mid] < arr[mid-1] && arr[mid] > arr[mid+1] {
				end = mid - 1
			}
		} else if mid-1 < 0 || mid+1 >= len(arr) {
			if arr[start] > arr[end] {
				return start
			} else {
				return end
			}
		}
	}
	return start
}
