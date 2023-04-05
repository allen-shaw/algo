package sort

import (
	"fmt"
	"math"
	"testing"
)

const (
	radix = 10 // 10进制
)

func RadixSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	digit := maxbits(arr)
	radixSort(arr, 0, len(arr)-1, digit)
}

func radixSort(arr []int, l, r, digit int) {
	if r-l < 2 {
		return
	}

	for i := 1; i <= digit; i++ { // 从第一位开始计算排序
		helper := make([]int, radix)
		tmp := make([]int, r-l+1)
		for j := l; j <= r; j++ {
			k := getBit(arr[j], i)
			helper[k]++
		}

		// 计算前缀和
		for i := 1; i < len(helper); i++ {
			helper[i] += helper[i-1]
		}

		// 根据前缀和出队，完成digit位排序
		for j := r; j >= l; j-- {
			// 从后往前遍历
			k := getBit(arr[j], i)
			pos := helper[k] - 1
			helper[k]--
			tmp[pos] = arr[j]
		}
		for i, v := range tmp {
			arr[l+i] = v
		}
	}
}

func maxbits(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return getDigit(max)
}

func getDigit(n int) int {
	k := 0
	for n > 0 {
		k++
		n /= 10
	}
	return k
}

func getBit(n int, k int) int {
	return n / int(math.Pow(radix, float64(k-1))) % 10
}

func TestGetDigit(t *testing.T) {
	fmt.Println(getDigit(8))
	fmt.Println(getDigit(10))
	fmt.Println(getDigit(190))
	fmt.Println(getDigit(9110))
}

func TestGetBit(t *testing.T) {
	n := 4035
	fmt.Println(getBit(n, 1))
	fmt.Println(getBit(n, 2))
	fmt.Println(getBit(n, 3))
	fmt.Println(getBit(n, 4))
	fmt.Println(getBit(n, 5))
}

func TestRadixSort(t *testing.T) {
	arr := []int{1, 9, 111, 25, 78, 39, 4, 19, 64, 77}
	RadixSort(arr)
	fmt.Println(arr)
}
