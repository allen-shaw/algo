package other

import "testing"

func TestMultiply(t *testing.T) {
	num1 := "123"
	num2 := "456"
	ans := multiply(num1, num2)
	t.Log(ans)	// 56088
}
