package _005_longest_palindromic_substring

//naive solution
//take two loop
//i = slow pointer, j = faster point.
// when i and jth element are equal, keep slowly moving the indexes to check if its palindrome
// maintain maxPalindrome and currentPalindrome

func longestPalindrome(s string) string {
	i, j := 0, len(s)-1
	//var longestPalindrome []byte
	b := []byte(s)
	var x, y int
	flag := 0
	for i < len(s)-1 {
		for j = len(s) - 1; j > i; j-- {
			if b[i] == b[j] {
				x = i
				y = j
				flag = 1
				// both pointer will move at the same pace
				for x <= y {
					//not palindrome
					if b[x] != b[y] {
						flag = 0
						break
					}
					x++
					y--
				}
				if flag == 1 {
					break
				}
			}
		}
		if flag == 1 {
			break
		}
		i++
	}
	println(x)
	println(y)
	//if there are no palindrome then return first character
	if i == len(s)-1 {
		return string(b[0])
	}
	return string(b[i : j+1])
}
