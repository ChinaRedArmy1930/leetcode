package main

func findKthPositive(arr []int, k int) int {
	newArr := make([]int, 0)
	for i := 1; i <= arr[len(arr)-1]+1; i++ {
		found := false
		for j := 0; j < len(arr); j++ {
			if i == arr[j] {
				found = true
				break
			}
		}
		if found {
			found = false
			continue
		}

		newArr = append(newArr, i)

	}
	//1,5,6,8,9,10

	if k <= len(newArr) {
		return newArr[k-1]
	}

	return newArr[len(newArr)-1] + (k - len(newArr))
}
