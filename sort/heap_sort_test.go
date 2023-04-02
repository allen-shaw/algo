package sort

import (
	"fmt"
	"testing"
)

// 堆排序，思路与冒泡排序/选择排序类似，每一轮将数组最大的值换到最后位置，进行n轮
// 相比冒泡排序/选择排序，堆排序改进了找最大值的过程，只需要log(N)时间复杂度

// 1. 先让数组变成一个大顶堆 N * log(N)
// 2. 将堆顶的元素和最后一个元素对调（相当于将堆顶元素（最大值）拿走，然后将堆底的那个元素补上它的空缺），
// 让那最后一个元素从顶上往下滑到恰当的位置（重新使堆最大化）。N * log(N)

func HeapSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	heapSort(arr, 0, len(arr)-1)
}

func heapSort(arr []int, l, r int) {
	// 1. 先让数组变成一个大顶堆
	size := 0
	for i := l; i <= r; i++ { // O(N)
		heapInsert(arr, i) // O(logN)
		size++
	}

	fmt.Println(size, arr)

	// 2. 将堆顶的元素和最后一个元素对调
	for size > 0 { // O(N)
		size--
		arr[0], arr[size] = arr[size], arr[0]
		heapify(arr, 0, size-1) // O(LogN)
	}
}

func heapInsert(arr []int, i int) {
	for {
		p := (i - 1) / 2
		if arr[i] <= arr[p] {
			// 比父节点小，停留原地，break
			break
		}
		// 比父节点大，与父节点交换，往堆顶升
		arr[i], arr[p] = arr[p], arr[i]
		i = p
	}
}

func heapify(arr []int, l, r int) {
	i := l
	for {
		lc := i*2 + 1
		if lc > r {
			// 没有孩子了，已经到最底，break
			break
		}
		//存在左孩子
		largest := lc
		if lc+1 <= r && arr[lc+1] > arr[lc] { //
			// 存在右孩子，且右孩子大于左孩子
			largest = lc + 1
		}

		if arr[i] > arr[largest] {
			// 父节点比子节点大，无需继续下沉，已经满足大顶堆，break
			break
		}
		// 交换父节点与比较大的子节点
		arr[i], arr[largest] = arr[largest], arr[i]
		i = largest
	}
}

func TestHeapInsert(t *testing.T) {
	arr := []int{6, 3, 7, 3, 9, 8, 1}
	// arr := []int{6, 3, 7, 3, 9}
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}
	fmt.Println(arr)
}

func TestHeapify(t *testing.T) {
	arr := []int{1, 7, 8, 3, 3, 6, 9}
	heapify(arr, 0, 5)
	fmt.Println(arr)
}

func TestHeapSort(t *testing.T) {
	arr := []int{6, 3, 7, 3, 9, 8, 1}
	HeapSort(arr)
	fmt.Println(arr)
}
