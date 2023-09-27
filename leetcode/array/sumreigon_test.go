package array

import (
	"fmt"
	"testing"
)

func TestSumRegion(t *testing.T) {
	// NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]);
	// numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
	// numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
	// numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)
	nm := NewNumMatrix([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5}})

	fmt.Println(nm.preSum)
	
	fmt.Println(nm.SumRegion(2, 1, 4, 3))
	fmt.Println(nm.SumRegion(1, 1, 2, 2))
	fmt.Println(nm.SumRegion(1, 2, 2, 4))
}
