package stack

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	str := "()[]{}"
	fmt.Println(isValid(str))
}
