package _049_Group_Anagrams

import (
	"sort"
)

// This method will group the anagrams together
func groupAnagrams(strs []string) [][]string {
	// map to add all the anagrams together with key being the sorted string
	m := make(map[string][]string)
	// loop array of strings
	for i := 0; i < len(strs); i++ {
		//sort the string
		key := sortCharacters(strs[i])
		// key is the sorted string and value is the original anagram string
		m[key] = append(m[key], strs[i])
	}

	result := make([][]string, 0, len(m))
	// loop in the map and fetch all the anagrams lists
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// this method will sort the characters of the string
func sortCharacters(s string) string {
	bs := []byte(s)
	//reference https://stackoverflow.com/questions/60768462/how-to-sort-bytes-in-string
	sort.Slice(bs, func(a, b int) bool {
		return bs[a] < bs[b]
	})
	return string(bs)
}
