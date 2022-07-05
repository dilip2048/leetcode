package _0006_Zigzag_Conversion

import "testing"

func TestZigZagConversionTest1(t *testing.T) {
	str := convert("paypalishiring", 4)
	if str != "pinalsigyahrpi" {
		t.Fail()
	}
}

func TestZigZagConversionTest2(t *testing.T) {
	str := convert("PAYPALISHIRING", 3)
	if str != "PAHNAPLSIIGYIR" {
		t.Fail()
	}
}

func TestZigZagConversionTest3(t *testing.T) {
	str := convert("A", 1)
	if str != "A" {
		t.Fail()
	}
}
