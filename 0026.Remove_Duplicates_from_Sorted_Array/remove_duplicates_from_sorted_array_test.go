package _026_Remove_Duplicates_from_Sorted_Array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	size := removeDuplicates(arr)
	if size != 5 {
		t.Fail()
	}
}
