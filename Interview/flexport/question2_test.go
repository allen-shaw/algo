package flexport

import (
	"fmt"
	"sort"
	"testing"
)

// 1.2 Unique Adjacent Elements in Array
// Given an integer array, produce an integer array where adjacent elements are not the same.
// Return any valid result, or an empty array if no solution exists.
// For example, given input [1,1,1,2,2,2], a possible output is [1,2,1,2,1,2].

type Pair struct {
	Num   int
	Count int
}

func uniqueAdjacentEleInArr(nums []int) []int {
	result := make([]int, len(nums))

	countMap := make(map[int]*Pair)
	for _, num := range nums {
		if p, ok := countMap[num]; ok {
			p.Count++
		} else {
			countMap[num] = &Pair{num, 1}
		}
	}

	list := make([]*Pair, 0, len(countMap))
	for _, p := range countMap {
		list = append(list, p)
	}

	// sort by count
	sort.Slice(list, func(i, j int) bool {
		return list[i].Count > list[j].Count
	})

	for i, idx := 0, 0; i < len(nums); i, idx = i+1, idx+2 {
		if idx >= len(nums) {
			idx = 1
		}

		n := list[0]
		n.Count--
		if n.Count == 0 {
			list = list[1:]
		}

		result[idx] = n.Num
		if idx > 0 && result[idx] == result[idx-1] {
			return nil
		}
		if idx < len(nums)-1 && result[idx] == result[idx+1] {
			return nil
		}
	}

	return result
}

func Test_uniqueAdjacentEleInArr(t *testing.T) {
	// nums := []int{1, 1, 1, 1, 2, 2, 3, 3}
	nums := []int{}
	out := uniqueAdjacentEleInArr(nums)
	fmt.Println(out)
}
