package _004_median_sorted_arrays

import "testing"

func TestMedian1(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	median := findMedian(nums1, nums2)
	if median != 2.5 {
		t.Fail()
	}
}

func TestMedian2(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{2, 3}
	median := findMedian(nums1, nums2)
	if median != 2 {
		t.Fail()
	}
}
