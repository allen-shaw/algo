package linklist

import "testing"

func TestAddTwoNumber(t *testing.T) {
	// l1 = [2,4,3], l2 = [5,6,4]
	l1 := buildList([]int{2, 4, 3})
	l2 := buildList([]int{5, 6, 4})
	printList(addTwoNumbers(l1, l2))
}

func TestAddTwoNumber2(t *testing.T) {
	l1 := buildList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := buildList([]int{9, 9, 9, 9})
	printList(addTwoNumbers(l1, l2))
}

func TestAddTwoNumber3(t *testing.T) {
	l1 := buildList([]int{9, 9})
	l2 := buildList([]int{9})
	printList(addTwoNumbers(l1, l2))
}
