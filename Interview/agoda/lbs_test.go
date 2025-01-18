package agoda

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Longest Bitonic Subsequence
// Given an array arr[] containing n positive integers,
// a subsequence of nums is called bitonic if it is first strictly increasing,
// then strictly decreasing. The task is to find the length of the longest bitonic subsequence.
// Note: A strictly increasing or a strictly decreasing sequence should not be considered as a bitonic sequence.
// Examples:
// Input: arr[]= [12, 11, 40, 5, 3, 1]
// Output: 5
// Explanation: The Longest Bitonic Subsequence is {12, 40, 5, 3, 1} which is of length 5.
// Input: arr[] = [80, 60, 30]
// Output: 0
// Explanation: There is no possible Bitonic Subsequence.

func longestBitonicSubsequence(arr []int) int {
	n := len(arr)

	dpInc := make([]int, n)
	dpDec := make([]int, n)
	for i := 0; i < n; i++ {
		dpInc[i] = 1
		dpDec[i] = 1
	}

	// increating subsequence
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				dpInc[i] = max(dpInc[j]+1, dpInc[i])
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if arr[j] < arr[i] {
				dpDec[i] = max(dpDec[j]+1, dpDec[i])
			}
		}
	}

	fmt.Println("increasing dp:", dpInc)
	fmt.Println("decreasing dp:", dpDec)
	lbs := 0
	for i := 0; i < n; i++ {
		if dpDec[i] == 1 || dpInc[i] == 1 {
			continue
		}
		lbs = max(dpInc[i]+dpDec[i]-1, lbs)
	}
	return lbs
}

func Test_longestBitonicSubsequence(t *testing.T) {
	testCases := []struct {
		arr    []int
		expect int
	}{
		{arr: []int{12, 11, 40, 5, 3, 1}, expect: 5},
		{arr: []int{80, 60, 30}, expect: 0},
	}

	for _, c := range testCases {
		res := longestBitonicSubsequence(c.arr)
		assert.Equal(t, c.expect, res)
	}
}
