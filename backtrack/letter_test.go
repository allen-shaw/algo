package backtrack

import (
	"fmt"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	digists := "2"
	ans := letterCombinations(digists)
	fmt.Println(ans)
	// ["ad","ae","af","bd","be","bf","cd","ce","cf"]
}

