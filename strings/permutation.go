package strings

import (
	"github.com/dllatas/algorithms/sorting"
)

func permUsingHash(s1, s2 string) bool {
	var size = len(s1)
	if size != len(s2) {
		return false
	}

	// this could be replaced by an array instead to save space
	var hash = make(map[rune]bool)

	for _, char := range s1 {
		_, found := hash[char]
		if !found {
			hash[char] = false
		}
	}

	for _, char := range s2 {
		_, found := hash[char]
		if !found {
			return false
		}
	}

	return true
}

func permUsingSort(s1, s2 string) bool {
	var size = len(s1)
	if size != len(s2) {
		return false
	}

	s1 = sorting.SortString(s1)
	s2 = sorting.SortString(s2)

	for i := 0; i < size; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
