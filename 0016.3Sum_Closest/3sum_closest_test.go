package _016_3Sum_Closest

import "testing"

func Test3SumCloset1(t *testing.T) {
	x := threeSumClosest([]int{-1, 2, 1, -4}, 1)
	if x != 2 {
		t.Fail()
	}
}

func Test3SumCloset2(t *testing.T) {
	x := threeSumClosest([]int{0, 0, 0}, 1)
	if x != 0 {
		t.Fail()
	}
}

func Test3SumCloset3(t *testing.T) {
	x := threeSumClosest([]int{-1, -2, -1, 8}, 1)
	if x != 5 {
		t.Fail()
	}
}
