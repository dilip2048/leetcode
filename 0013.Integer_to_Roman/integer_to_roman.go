package _013_Integer_to_Roman

import (
	"fmt"
	"sort"
)

// This method converts integer to roman number
func intToRoman(num int) string {
	m := make(map[int]string, 7)
	m[1000] = "M"
	m[900] = "CM"
	m[500] = "D"
	m[400] = "CD"
	m[100] = "C"
	m[90] = "XC"
	m[50] = "L"
	m[40] = "XL"
	m[10] = "X"
	m[9] = "IX"
	m[5] = "V"
	m[4] = "IV"
	m[1] = "I"
	keys := make([]int, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var result string
	i := len(keys) - 1
	for num > 0 {
		if num >= keys[i] {
			result = fmt.Sprintf("%s%s", result, m[keys[i]])
			num -= keys[i]
		} else {
			i--
		}
	}
	return result
}
