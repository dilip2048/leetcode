package twosum

import "testing"

func TestTwoSumBaseCase(t *testing.T) {
	a := []int{-1, 0, 1, 2, -1, -4}
	x := threeSum(a)
	//todo: compare the threeSum() output with the expected data
	if x == nil {
		t.Fail()
	}
}
