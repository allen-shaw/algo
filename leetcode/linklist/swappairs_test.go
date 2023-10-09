package linklist

import "testing"

func TestSwapPair(t *testing.T) {
	head := buildList([]int{1, 2, 3, 4})
	head = swapPairs(head)
	printList(head)
}
