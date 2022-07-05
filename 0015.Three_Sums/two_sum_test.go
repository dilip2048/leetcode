package twosum

import "testing"

func TestTwoSumBaseCase(t *testing.T) {
	var a []int
	x := threeSum(a, 6)
	if x != nil {
		t.Fatal("test failed")
	}
}

func TestTwoSum1(t *testing.T) {
	a := []int{1, 2, 3, 4}
	x := threeSum(a, 6)
	if x[0] != 1 || x[1] != 3 {
		t.Fatal("test failed")
	}
}

func TestTwoSum2(t *testing.T) {
	a := []int{3, 3}
	x := threeSum(a, 6)
	if x[0] != 0 || x[1] != 1 {
		t.Fatal("test failed")
	}
}
