package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	less, greater := partition(arr, l, r)
	quickSort(arr, l, less)
	quickSort(arr, greater, r)
}

func partition(arr []int, l, r int) (int, int) {
	k := l + rand.Intn(r-l)

	less, greater := l-1, r+1

	for i := l; i < greater; {
		if arr[i] < arr[k] {
			less++
			arr[i], arr[less] = arr[less], arr[i]
			i++
		} else if arr[i] == arr[k] {
			i++
		} else {
			greater--
			arr[i], arr[greater] = arr[greater], arr[i]
		}
	}

	return less, greater
}

func TestQuickSort(t *testing.T) {
	arr := []int{1, 9, 11, 25, 78, 3, 4, 19, 64, 77}
	QuickSort(arr)
	fmt.Println(arr)
}
