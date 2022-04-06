package _005_longest_palindromic_substring

import (
	"strings"
	"testing"
)

func TestLongestPalindromeBaseCase1(t *testing.T) {
	s := "a"
	str := longestPalindrome(s)
	if !strings.EqualFold("a", str) {
		t.Fail()
	}
}

func TestLongestPalindromeBaseCase2(t *testing.T) {
	s := "abc"
	str := longestPalindrome(s)
	if !strings.EqualFold("a", str) {
		t.Fail()
	}
}

func TestLongestPalindrome1(t *testing.T) {
	s := "babad"
	str := longestPalindrome(s)
	if !strings.EqualFold("bab", str) {
		t.Fail()
	}
}

func TestLongestPalindrome2(t *testing.T) {
	s := "aacabdkacaa"
	str := longestPalindrome(s)
	if !strings.EqualFold("bab", str) {
		t.Fail()
	}
}

func TestLongestPalindrome3(t *testing.T) {
	s := "ac"
	str := longestPalindrome(s)
	if !strings.EqualFold("a", str) {
		t.Fail()
	}
}
