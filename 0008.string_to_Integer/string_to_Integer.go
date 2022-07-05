package _008_string_to_Integer

import "math"

func Atoi(s string) int {
	l := len(s)
	j := 0
	//flag for positive and negative number
	sign := 1
	//delete all the spaces first
	for j < l && s[j] == 32 {
		j++
	}
	// check for sign
	if j < l && (s[j] == '+' || s[j] == '-') {
		if s[j] == '-' {
			sign = -1
		}
		j++
	}
	num := 0
	// flag to check if the leading string is integer
	var flag bool
	for i := j; i < l; i++ {
		if s[i] >= 48 && s[i] <= 57 {
			num = num*10 + int(s[i]-'0')
			// check for sign and check if it exceed max value of int32
			if num > math.MaxInt32 && sign == -1 {
				return math.MinInt32
			} else if num > math.MaxInt32 && sign == 1 {
				return math.MaxInt32
			}
			flag = true
		} else {
			break
		}
	}
	if flag {
		return num * sign
	}
	return 0
}
