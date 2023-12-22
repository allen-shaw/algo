package sort

import (
	gosort "sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	ss := make([]string, len(nums))

	for i, n := range nums {
		ss[i] = strconv.Itoa(n)
	}

	gosort.Slice(ss, func(i, j int) bool {
		return ss[i]+ss[j] > ss[j]+ss[i]
	})

	i := 0
	for ; i < len(ss)-1; i++ {
		if ss[i] != "0" {
			break
		}
	}

	return strings.Join(ss[i:], "")

}
