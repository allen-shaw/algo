package other

import (
	"fmt"
	"testing"
)

// 长度为n-1的整形数组，数字的范围在1到n，找其中一个缺失的数字
func getLoseNum1(nums []int) int {
	ans := 0
	for _, n := range nums {
		ans ^= n
	}
	for i := 0; i <= len(nums); i++ {
		ans ^= i
	}
	return ans
}
func TestGetLoseNum1(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 6, 7}
	ans := getLoseNum1(nums)
	fmt.Println(ans)
}

func findSingleNumber(nums []int) (int, int) {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}

	temp := xor
	k := 0
	for temp != 0 {
		if temp&1 == 1 {
			break
		} else {
			k++
			temp = temp >> 1
		}
	}

	x, y := 0, 0
	for _, n := range nums {
		if (n >> k) == 1 {
			x ^= n
		} else {
			y ^= n
		}
	}

	return x, y
}

func TestFindSingle(t *testing.T) {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(findSingleNumber(nums))
}

// func findLoseNum2(nums []int) {

// }

func TestXOR(t *testing.T) {
	a := -820
	fmt.Printf("a:%b\n", a)
	k := 0
	for a != 0 {
		if a&1 == 1 {
			break
		} else {
			k++
			a = a >> 1
		}
	}

	fmt.Println(k)

	// c :=
}
