package sort

// 从0开始，和隔壁数字比较，谁大谁往右移动
// 一轮过后，最大的数被推到右边
// 进行n-1轮后，完成排序
func BubbleSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// 从0到arr.end进行一轮冒泡，这一轮后，最大值到最右
	for end := len(arr) - 1; end > 0; end-- {
		// 对0到end的数进行一轮冒泡
		bubble(arr, 0, end)
	}
}

func bubble(arr []int, begin, end int) {
	// 这里用i < end，而不是 i <= end是因为下面用到arr[i]和arr[i+1]
	// 因此i到arr倒数第二位，实际上已经比较玩所有有数据了
	for i := begin; i < end; i++ {
		// 将更大的值推到右边
		if arr[i] > arr[i+1] {
			swap(arr, i, i+1)
		}
	}
}
