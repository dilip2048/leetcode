package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	l := len(s)
	//base case
	if l == 1 {
		return true
	}
	var low, high int
	//odd case
	if l%2 == 1 {
		low = l/2 - 1
		high = l/2 + 1
	} else {
		low = l/2 - 1
		high = l / 2
	}
	for low >= 0 && high <= l-1 {
		if s[low] != s[high] {
			return false
		}
		low--
		high++
	}
	return true
}