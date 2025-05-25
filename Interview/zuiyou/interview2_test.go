package zuiyou

import (
	"fmt"
	"testing"
)

// [[1,1],2,[1,1]]
func depthSum(nestedList string) int {
	stack := make([]byte, 0)
	sum := 0

	for i := 0; i < len(nestedList); {
		c := nestedList[i]
		if c == '[' {
			stack = append(stack, c)
			i++
		} else if c == ']' {
			stack = stack[:len(stack)-1]
			i++
		} else if c >= '0' && c <= '9' {
			num := 0
			for nestedList[i] >= '0' && nestedList[i] <= '9' {
				num = num * 10 + int(nestedList[i]-'0')
				i++
			}
			sum += num * len(stack)
			fmt.Println(i, string(nestedList[i]), sum)
		} else {
			i++
		}
	}

	return sum
}

func Test_depthSum(t *testing.T) {
	fmt.Println(depthSum("[[1,1],2,[1,1]]"))
	fmt.Println(depthSum("[1,[4,[6]]]"))
}
