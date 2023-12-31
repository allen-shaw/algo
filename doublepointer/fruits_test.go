package doublepointer

import (
	"fmt"
	"testing"
)

func TestTotalFruits(t *testing.T) {
	// fruits := []int{1, 2, 1}
	// fruits := []int{0, 1, 2, 2}
	// fruits := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fruits := []int{1, 0, 1, 4, 1, 4, 1, 2, 3}
	ans := totalFruit(fruits)
	fmt.Println(ans)
}
