package _018_4Sum

import "testing"

func TestTwoSumBaseCase(t *testing.T) {
	a := []int{2, 2, 2, 2, 2, 2}
	x := fourSum(a, 8)
	//todo: compare the threeSum() output with the expected data
	if x == nil {
		t.Fail()
	}
}
