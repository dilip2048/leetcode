package _053_Maximum_Subarray

import "testing"

func TestMaximumSubarray1(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	x := maxSubArray(arr)
	if x != 6 {
		t.Fail()
	}
}

func TestMaximumSubarray2(t *testing.T) {
	arr := []int{5, 4, -1, 7, 8}
	x := maxSubArray(arr)
	if x != 23 {
		t.Fail()
	}
}
