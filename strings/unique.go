package strings

import (
	"sort"

	"github.com/dllatas/algorithms/sorting"
)

func uniqueUsingHash(input string) bool {
	var buffer = make(map[rune]bool)

	for _, char := range input {
		_, found := buffer[char]
		if !found {
			buffer[char] = true
			continue
		}
		return false
	}

	return true
}

func unique(input string) bool {
	var size = len(input)

	for i := 0; i < size; i++ {
		for j := i + 1; j <= size-1; j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}

	return true
}

func uniqueUsingBitVector(input string) bool {
	var size = len(input)
	const asciiSize = 128
	var bit [asciiSize]int = [asciiSize]int{}

	for i := 0; i < size; i++ {
		var index = int(input[i])
		if bit[index] == 1 {
			return false
		}
		bit[index] = 1
	}

	return true
}

func uniqueUsingSort(input string) bool {
	s := []rune(input)
	sort.Sort(sorting.SortRunes(s))

	var size = len(input)

	for i := 0; i < size-1; i++ {
		if s[i] == s[i+1] {
			return false
		}
	}

	return true
}
