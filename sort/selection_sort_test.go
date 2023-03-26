package sort

// 遍历一遍数组，选最小值放到第一位
func SelectionSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for i := range arr {
		// 找到i到arr.end访问内最小值的index
		minIndex := min(arr, i)
		// 与arr[i]交换
		swap(arr, i, minIndex)
	}
}

func min(arr []int, beginIndex int) int {
	minValue := arr[beginIndex]
	index := beginIndex
	for i := beginIndex; i < len(arr); i++ {
		if arr[i] < minValue {
			minValue = arr[i]
			index = i
		}
	}
	return index
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
