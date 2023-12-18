package doublepointer

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	link := buildLinkedList([]int{1, 2, 3, 3, 4, 4, 5})
	deleteDuplicates2(link)

	t.Log(link.String())
}

func TestDeleteDuplicatesR(t *testing.T) {
	link := buildLinkedList([]int{1, 2, 3, 3, 4, 4, 5})
	deleteDuplicates2R(link)

	t.Log(link.String())
}

func TestDeleteDuplicatesP(t *testing.T) {
	link := buildLinkedList([]int{1, 2, 3, 3, 4, 4, 5})
	deleteDuplicatesP(link)

	t.Log(link.String())
}
