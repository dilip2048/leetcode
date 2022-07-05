package _008_string_to_Integer

import "testing"

func TestAtoiBaseCase1(t *testing.T) {
	x := Atoi("")
	if x != 0 {
		t.Fail()
	}
}

func TestAtoiBaseCase2(t *testing.T) {
	x := Atoi("   ")
	if x != 0 {
		t.Fail()
	}
}

func TestAtoiEdgeCase1(t *testing.T) {
	x := Atoi("-91283472332")
	if x != -2147483648 {
		t.Fail()
	}
}

func TestAtoiEdgeCase2(t *testing.T) {
	x := Atoi("-2147483647")
	if x != -2147483647 {
		t.Fail()
	}
}

func TestAtoiEdgeCase3(t *testing.T) {
	x := Atoi("21474836460")
	if x != 2147483647 {
		t.Fail()
	}
}

func TestAtoi1(t *testing.T) {
	x := Atoi("   -42")
	if x != -42 {
		t.Fail()
	}
}

func TestAtoi2(t *testing.T) {
	x := Atoi("4193 with words")
	if x != 4193 {
		t.Fail()
	}
}

func TestAtoi3(t *testing.T) {
	x := Atoi("42")
	if x != 42 {
		t.Fail()
	}
}
