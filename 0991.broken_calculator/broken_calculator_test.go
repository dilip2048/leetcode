package _991_broken_calculator

import "testing"

func TestBrokenCalcBaseCase(t *testing.T) {
	counter := brokenCalc(1, 1)
	if counter != 0 {
		t.Fail()
	}
}

func TestBrokenCalcCase1(t *testing.T) {
	counter := brokenCalc(2, 3)
	if counter != 2 {
		t.Fail()
	}
}

func TestBrokenCalcCase2(t *testing.T) {
	counter := brokenCalc(5, 8)
	if counter != 2 {
		t.Fail()
	}
}

func TestBrokenCalcCase3(t *testing.T) {
	counter := brokenCalc(3, 10)
	if counter != 3 {
		t.Fail()
	}
}
