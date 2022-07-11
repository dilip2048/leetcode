package _049_Group_Anagrams

import (
	"testing"
)

func TestGroupAnagramBaseCase1(t *testing.T) {
	arr := []string{""}
	result := [][]string{{""}}
	rLen := len(result)
	anagrams := groupAnagrams(arr)
	for i := 0; i < rLen; i++ {
		for j := 0; j < len(result[i]); i++ {
			if result[i][j] != anagrams[i][j] {
				t.Fail()
			}
		}
	}
}

func TestGroupAnagramBaseCase2(t *testing.T) {
	arr := []string{"a"}
	result := [][]string{{"a"}}
	rLen := len(result)
	anagrams := groupAnagrams(arr)
	for i := 0; i < rLen; i++ {
		for j := 0; j < len(result[i]); i++ {
			if result[i][j] != anagrams[i][j] {
				t.Fail()
			}
		}
	}
}

func TestGroupAnagram1(t *testing.T) {
	arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}
	rLen := len(result)
	anagrams := groupAnagrams(arr)
	for i := 0; i < rLen; i++ {
		for j := 0; j < len(result[i]); i++ {
			if result[i][j] != anagrams[i][j] {
				t.Fail()
			}
		}
	}
}
