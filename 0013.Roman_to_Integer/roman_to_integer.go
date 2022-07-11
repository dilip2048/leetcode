package _013_Integer_to_Roman

func romanToInt(s string) int {
	m := make(map[string]int, 7)
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000
	var sum int
	var i int
	for i = 0; i < len(s); i++ {
		if i+1 < len(s) && m[string(s[i+1])] > m[string(s[i])] {
			sum -= m[string(s[i])]
		} else {
			sum += m[string(s[i])]
		}
	}
	return sum
}
