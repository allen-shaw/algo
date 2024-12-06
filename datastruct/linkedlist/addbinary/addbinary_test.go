package addbinary

import (
	"fmt"
	"testing"
)

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	l := max(la, lb) + 1

	ans := make([]byte, l)
	carry := byte(0)

	i, pa, pb := l-1, la-1, lb-1
	for pa >= 0 || pb >= 0 {
		val := carry
		if pa >= 0 {
			val += a[pa] - '0'
			pa--
		}
		if pb >= 0 {
			val += b[pb] - '0'
			pb--
		}
		carry = val / 2
		val = val % 2
		ans[i] = val + '0'
		i--
	}

	if carry == 0 {
		ans = ans[1:]
	} else {
		ans[0] = carry + '0'
	}

	return string(ans)
}

func Test_addBinary(t *testing.T) {
	a := "1010" // 10101
	b := "1011"
	out := addBinary(a, b)
	fmt.Println("out:", out)
}
