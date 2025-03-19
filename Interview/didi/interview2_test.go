package didi

import (
	"errors"
	"math"
)

// StringToInt
func StringToInt(str string) (int, error) {
	var res int
	flag := 1
	if str[0] == '-' {
		flag = -1
		str = str[1:]
	}

	for _, n := range str {
		if n >= '0' && n <= '9' {
			if res > math.MaxInt/10 {
				return 0, errors.New("overflow")
			}
			val := int(n - '0')
			res = res*10 + val
		} else {
			return 0, errors.New("invalid string")
		}
	}

	return res * flag, nil
}
