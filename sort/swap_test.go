package sort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	a := 100
	b := 200

	fmt.Println(a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a, b)
}
