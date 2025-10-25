package avenir

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := []int{0,1,2,3,4,5,6}
	fmt.Println(arr[1:3])
	fmt.Println(arr[3:])
}
