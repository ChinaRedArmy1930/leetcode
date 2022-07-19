package main

var a = 101

//方案1
func corpFlightBookings(bookings [][]int, n int) []int {
	ret := make([]int, n+1)

	for i := 0; i < len(bookings); i++ {
		for j := bookings[i][0]; j <= bookings[i][1]; j++ {
			ret[j] += bookings[i][2]
		}
	}

	return ret[1:len(ret)]
}
