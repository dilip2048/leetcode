package _026_Remove_Duplicates_from_Sorted_Array

import "testing"

func TestRemoveElement(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	size := removeElement(arr, 2)
	if size != 8 {
		t.Fail()
	}
}
