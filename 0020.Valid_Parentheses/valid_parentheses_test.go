package _020_Valid_Parentheses

import "testing"

func TestValidParanthesis1(t *testing.T) {
	b := isValid("()")
	if b != true {
		t.Fail()
	}
}

func TestValidParanthesis2(t *testing.T) {
	b := isValid("()[]{}")
	if b != true {
		t.Fail()
	}
}

func TestValidParanthesis3(t *testing.T) {
	b := isValid("(")
	if b != false {
		t.Fail()
	}
}

func TestValidParanthesis4(t *testing.T) {
	b := isValid("{{}[][[[]]]}")
	if b != true {
		t.Fail()
	}
}

func TestValidParanthesis6(t *testing.T) {
	b := isValid(")}{}")
	if b != false {
		t.Fail()
	}
}
