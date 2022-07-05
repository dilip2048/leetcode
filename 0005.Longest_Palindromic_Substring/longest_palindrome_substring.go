package _005_Longest_Palindromic_Substring

func NaiveLongestPalindrome(s string) string {
	// all single character string is palindrome
	maxSubStringLength := 1
	startingIndex := 0
	//this loop start from starting index of the substring
	for i := 0; i < len(s); i++ {
		// this loop is to point the end index of the substring
		for j := i; j < len(s); j++ {
			flag := 1
			// this loop is to check if the substring is palindrome from ith of jth index of the string
			for k := 0; k < (j-i+1)/2; k++ {
				if s[i+k] != s[j-k] {
					flag = 0
				}
			}
			if flag == 1 && maxSubStringLength < (j-i+1) {
				maxSubStringLength = j - i + 1
				startingIndex = i
			}
		}
	}
	palindromeString := s[startingIndex : startingIndex+maxSubStringLength]
	return palindromeString
}

// OptimizedLongestPalindrome :Use expand Around the center approach to solve the problem.
func OptimizedLongestPalindrome(s string) string {
	//find palindrome for even length palindrome substring
	maxStringLen := 1
	start := 0
	for i := 1; i < len(s); i++ {
		low := i - 1
		high := i
		for low >= 0 && high < len(s) && s[low] == s[high] {
			if high-low+1 > maxStringLen {
				start = low
				maxStringLen = high - low + 1
			}
			//expand arround the centre
			low--
			high++
		}
		// check for palindrome for odd length substring
		low = i - 1
		high = i + 1
		for low >= 0 && high < len(s) && s[low] == s[high] {
			if high-low+1 > maxStringLen {
				start = low
				maxStringLen = high - low + 1
			}
			low--
			high++
		}
	}
	return s[start : start+maxStringLen]
}
