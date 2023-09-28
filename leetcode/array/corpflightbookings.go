package array

func corpFlightBookings(bookings [][]int, n int) []int {
	sitediff := make([]int, 10001)

	for _, b := range bookings {
		diffAdd(sitediff, b[0], b[1], b[2])
	}

	return diffResume(sitediff, n)
}

func diffAdd(arr []int, first, last, n int) {
	arr[first] += n
	arr[last+1] -= n
}

func diffResume(arr []int, n int) []int {
	res := make([]int, n+2)
	for i := 0; i < n+1; i++ {
		res[i+1] = res[i] + arr[i]
	}
	return res[2:]
}
