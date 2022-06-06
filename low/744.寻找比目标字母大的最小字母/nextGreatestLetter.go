package main

func nextGreatestLetter(letters []byte, target byte) byte {
	start := 0
	end := len(letters) - 1

	for start < end {
		mid := start + (end-start)/2

		if letters[mid] > target {
			end = mid
		} else {
			start = mid + 1
		}
	}

	if letters[end] <= target {
		return letters[0]
	}
	return letters[end]
}
