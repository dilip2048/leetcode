package _005_longest_palindromic_substring

import (
	"strings"
	"testing"
)

func TestLongestPalindromeBaseCase1(t *testing.T) {
	s := "aa"
	str := OptimizedLongestPalindrome(s)
	if !strings.EqualFold("aa", str) {
		t.Fail()
	}
}

func TestLongestPalindromeBaseCase2(t *testing.T) {
	s := "abc"
	str := NaiveLongestPalindrome(s)
	if !strings.EqualFold("a", str) {
		t.Fail()
	}
}

func TestLongestPalindrome1(t *testing.T) {
	s := "babad"
	str := NaiveLongestPalindrome(s)
	if !strings.EqualFold("bab", str) {
		t.Fail()
	}
}

func TestLongestPalindrome2(t *testing.T) {
	s := "aacaabdkacaa"
	str := NaiveLongestPalindrome(s)
	if !strings.EqualFold("aacaa", str) {
		t.Fail()
	}
}

func TestLongestPalindrome3(t *testing.T) {
	s := "ac"
	str := OptimizedLongestPalindrome(s)
	if !strings.EqualFold("a", str) {
		t.Fail()
	}
}
