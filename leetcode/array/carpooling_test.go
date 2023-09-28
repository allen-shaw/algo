package array

import (
	"fmt"
	"testing"
)

func TestCarPooling(t *testing.T) {
	trips := [][]int{{2, 1, 5}, {3, 3, 7}}
	capacity := 5

	fmt.Println(carPooling(trips, capacity))
}
