package _0006_Zigzag_Conversion

func convert(s string, numRows int) string {
	a := make([][]uint8, numRows)
	for i := range a {
		a[i] = make([]uint8, 0)
	}
	var i int
	for i < len(s) {
		for index := 0; index < numRows && i < len(s); index++ {
			a[index] = append(a[index], s[i])
			i++
		}
		for index := numRows - 2; index >= 1 && i < len(s); index-- {
			a[index] = append(a[index], s[i])
			i++
		}
	}
	for index := 1; index < numRows; index++ {
		a[0] = append(a[0], a[index]...)
	}
	return string(a[0])
}
