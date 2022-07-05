package twosum

// This is a sliding window and hashing problem
// Approach reference from https://www.enjoyalgorithms.com/blog/longest-substring-without-repeating-characters

func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]string)
	longestSubStrLen := 0
	subStrLen := 0
	i, j := 0, 0
	for i < len(s) && j < len(s) {
		if _, ok := m[s[j]]; !ok {
			subStrLen = j - i + 1
			if subStrLen > longestSubStrLen {
				longestSubStrLen = subStrLen
			}
			m[s[j]] = "visited"
			j = j + 1
		} else {
			// delete the first occurrence of the character
			delete(m, s[i])
			i = i + 1
		}
	}
	return longestSubStrLen
}
