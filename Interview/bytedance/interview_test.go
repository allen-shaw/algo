package bytedance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// a - b
func subString(a, b string) string {
	if (len(a) < len(b)) || (len(a) == len(b) && a < b) {
		return "-" + subString(b, a)
	}

	result := make([]byte, 0)
	flag := 0
	i, j := len(a)-1, len(b)-1
	for i >= 0 && j >= 0 {
		ans := ctoi(a[i]) - ctoi(b[j]) - flag
		flag = 0
		if ans < 0 {
			ans += 10
			flag = 1
		}
		result = append(result, itoc(ans))
		i--
		j--
	}

	for i >= 0 {
		ans := ctoi(a[i]) - flag
		flag = 0
		if ans < 0 {
			ans += 10
			flag = 1
		}
		result = append(result, itoc(ans))
		i--
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	idx := 0
	for idx < len(result)-1 {
		if result[idx] != '0' {
			break
		}
		idx++
	}

	return string(result[idx:])
}

func ctoi(c byte) int {
	return int(c - '0')
}

func itoc(i int) byte {
	return byte(i) + '0'
}

func Test_subString(t *testing.T) {
	testCases := []struct {
		a, b   string
		expect string
	}{
		{a: "11", b: "123", expect: "-112"},
		{a: "1000000", b: "9", expect: "999991"},
		{a: "1000000", b: "999999", expect: "1"},
		{a: "456", b: "77", expect: "379"},
	}

	for _, tc := range testCases {
		ans := subString(tc.a, tc.b)
		assert.Equal(t, tc.expect, ans)
	}
}
