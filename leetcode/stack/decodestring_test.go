package stack

import (
	"fmt"
	"testing"
)

func TestIsNum(t *testing.T) {
	a := byte('a')
	fmt.Println(a, " is a num ", isNum(a))

	b := byte('1')
	fmt.Println(b, " is a num ", isNum(b))
}

func TestCToNum(t *testing.T) {
	a := byte('2')
	fmt.Println(a, " to num ", ctoNum(a))
}

func TestDecodeString(t *testing.T) {
	s := "3[a]2[bc]"
	ans := decodeString(s)
	fmt.Println(ans)
	// aaabcbc
}

func TestDecodeString2(t *testing.T) {
	s := "3[a2[c]]"
	ans := decodeString(s)
	fmt.Println(ans)
	// accaccacc
}

func TestDecodeString3(t *testing.T) {
	s := "2[abc]3[cd]ef"
	ans := decodeString(s)
	fmt.Println(ans)
	// abcabccdcdcdef
}
