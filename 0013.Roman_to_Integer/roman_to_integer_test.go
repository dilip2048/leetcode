package _013_Integer_to_Roman

import "testing"

func TestRomanToInteger1(t *testing.T) {
	integer := romanToInt("III")
	if integer != 3 {
		t.Fail()
	}
}

func TestRomanToInteger2(t *testing.T) {
	integer := romanToInt("LVIII")
	if integer != 58 {
		t.Fail()
	}
}
