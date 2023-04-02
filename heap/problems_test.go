package heap

import (
	"fmt"
	"testing"
)

// 修改大顶堆中任意一个元素（坐标为i)，恢复为大顶堆
func RecoverHeap(arr []int, i int) {
	heapInsert(arr, i, arr[i])
	doHeapify(arr, i, len(arr)-1)
}

func TestRecoverHeap(t *testing.T) {
	// {8, 7, 4, 3, 1}
	arr := []int{8, 7, 10, 3, 1}

	// 变大
	RecoverHeap(arr, 2)
	fmt.Println(arr)

	// 不变
	arr2 := []int{8, 7, 4, 3, 1}
	RecoverHeap(arr2, 2)
	fmt.Println(arr2)

	// // 变小
	arr3 := []int{8, 0, 4, 3, 1}
	RecoverHeap(arr3, 1)
	fmt.Println(arr3)
}

func doHeapify(arr []int, l, r int) {
	i := l
	for {
		lc := 2*i + 1
		largest := lc
		if lc >= r {
			// 没有子节点
			break
		}
		if lc+1 <= r && arr[lc+1] > arr[lc] {
			// 有右孩子，且右孩子大于左孩子
			largest = lc + 1
		}
		// 父节点与大的孩子交换
		arr[i], arr[largest] = arr[largest], arr[i]
		i = largest
	}
}

// 已知一个几乎有序的数组，把数组排好顺序的话，每个元素移动的距离一定不超过k，并且k相对于数组长度来说比较小
// 时间复杂度O(nlogk)
// max代表最大距离
func SortAlmostOrdered(arr []int, max int) {
	// 建立一个大小为max的小顶堆
	// 从0-max入堆，此时堆顶是整个数组的最小值
	// 位置大于max的数不可能是最小值，假设max+k是最小值，则移动距离为max+k,与题干矛盾
	heap := NewHeap(max)
	k := 0

	for _, v := range arr {
		heap.Insert(v)
		if heap.Size() == max {
			min, _ := heap.Pop()
			arr[k] = min
			k++
		}
	}

	for !heap.isEmpty() {
		min, _ := heap.Pop()
		arr[k] = min
	}
}

func TestSortAlmostOrdered(t *testing.T) {
	arr := []int{3, 2, 6, 9, 12, 11, 10}
	SortAlmostOrdered(arr, 2)
	fmt.Println(arr)
}
