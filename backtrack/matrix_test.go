package backtrack

import "testing"

func TestMatrix(t *testing.T) {
	ans := NewMatrix[int](3, 0)
	t.Log(ans)
}

func TestMatrix2(t *testing.T) {
	m := NewMatrix[byte](4, '.')
	ans := make([]string, 0, len(m))

	for i := 0; i < len(m); i++ {
		ans = append(ans, string(m[i]))
	}

	t.Log(ans, len(ans))
}
