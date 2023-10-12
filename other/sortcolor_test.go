package other

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func TestSortColors2(t *testing.T) {
	nums := []int{1, 2, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func TestSortColors3(t *testing.T) {
	nums := []int{2}
	sortColors(nums)
	fmt.Println(nums)
}
