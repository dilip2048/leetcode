package _011_Container_With_Most_Water

import "testing"

func TestTwoSumBaseCase1(t *testing.T) {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := maxArea(a)
	if area != 49 {
		t.Fail()
	}
}

func TestTwoSumBaseCase2(t *testing.T) {
	a := []int{1, 1}
	area := maxArea(a)
	if area != 1 {
		t.Fail()
	}
}
