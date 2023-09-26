package array

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	arr1 := []int{1, 1, 2}
	n1 := removeDuplicates(arr1)
	fmt.Println(n1)

	arr2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n2 := removeDuplicates(arr2)

	fmt.Println(n2)
}
