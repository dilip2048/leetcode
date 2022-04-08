package main

import (
	"testing"
)

func TestReverseIntBaseCase(t *testing.T) {
	x := reverse(0)
	if x != 0 {
		t.Fail()
	}
}

func TestReverseIntEdgeCase(t *testing.T) {
	x := reverse(-123)
	if x != -321 {
		t.Fail()
	}
}

func TestReverseIntCase1(t *testing.T) {
	x := reverse(123)
	if x != 321 {
		t.Fail()
	}
}
