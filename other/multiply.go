package other

import (
	"strconv"
)

func multiply(num1 string, num2 string) string {
	res := make([]int, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			n1, n2 := int(num1[i]-'0'), int(num2[j]-'0')
			mul := n1*n2 + res[i+j+1]
			res[i+j+1] = mul % 10
			res[i+j] += mul / 10
		}
	}

	k := 0
	for ; k < len(res); k++ {
		if res[k] != 0 {
			break
		}
	}

	s := ""
	for i := k; i < len(res); i++ {
		s += strconv.Itoa(res[i])
	}

	return s
}
