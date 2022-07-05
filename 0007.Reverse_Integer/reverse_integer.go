package main

import (
	"log"
	"math"
	"strconv"
)

func reverse(x int) int {
	temp := x
	// take positive of already negative
	if x < 0 {
		temp *= -1
	}
	// convert the number to string
	s := strconv.Itoa(temp)
	// save converted string into byte array
	r := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	// convert the reversed bytes to integer
	reverseInt, err := strconv.Atoi(string(r))
	if err != nil {
		log.Fatal(err)
	}
	// check if the reversed number is greater then int32
	if reverseInt >= math.MaxInt32 {
		return 0
	}
	//add negation at the end if the number is negative
	if x < 0 {
		return reverseInt * -1
	}
	return reverseInt
}
