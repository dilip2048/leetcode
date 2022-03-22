package twosum

import (
	"testing"
)

func TestFindLongestSubstring1(t *testing.T) {
	str := "pwwbew"
	s := lengthOfLongestSubstring(str)
	if s != 3 {
		t.Fatal("test failed")
	}
}

func TestFindLongestSubstring2(t *testing.T) {
	str := "abcabcbb"
	s := lengthOfLongestSubstring(str)
	if s != 3 {
		t.Fatal("test failed")
	}
}

func TestFindLongestSubstring3(t *testing.T) {
	str := "dvdf"
	s := lengthOfLongestSubstring(str)
	if s != 3 {
		t.Fatal("test failed")
	}
}

func TestFindLongestSubstring4(t *testing.T) {
	str := " "
	s := lengthOfLongestSubstring(str)
	if s != 1 {
		t.Fatal("test failed")
	}
}
