package array

func carPooling(trips [][]int, capacity int) bool {
	tripsdiff := make([]int, 1001)

	for _, t := range trips {
		tripsAdd(tripsdiff, t[1], t[2], t[0])
	}

	res := tripResume(tripsdiff, len(tripsdiff)-1)
	for _, n := range res {
		if n > capacity {
			return false
		}
	}
	return true
}

func tripsAdd(arr []int, l, r, n int) {
	arr[l] += n
	arr[r] -= n
}

func tripResume(arr []int, n int) []int {
	res := make([]int, n+1)
	for i := 0; i < n; i++ {
		res[i+1] = res[i] + arr[i]
	}

	return res[1:]
}
