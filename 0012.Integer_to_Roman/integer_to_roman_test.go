package _012_Integer_to_Roman

import "testing"

func TestIntegerToRoman1(t *testing.T) {
	roman := intToRoman(3)
	if roman != "III" {
		t.Fail()
	}
}

func TestIntegerToRoman2(t *testing.T) {
	roman := intToRoman(58)
	if roman != "LVIII" {
		t.Fail()
	}
}

func TestIntegerToRoman3(t *testing.T) {
	roman := intToRoman(1994)
	if roman != "MCMXCIV" {
		t.Fail()
	}
}
