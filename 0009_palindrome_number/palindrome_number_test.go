package _0009_palindrome_number

import (
	"testing"
)

func TestPalindromeNumberBaseCase(t *testing.T) {
	b := isPalindrome(0)
	if b != true {
		t.Fail()
	}
}

func TestPalindromeNumberFalseCase(t *testing.T) {
	b := isPalindrome(123)
	if b != false {
		t.Fail()
	}
}

func TestPalindromeNumbereTrueCase(t *testing.T) {
	b := isPalindrome(121)
	if b != true {
		t.Fail()
	}
}
