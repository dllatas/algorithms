package unique_chars

import (
	"sort"
	"strings"
)

func uniqueSort(input string) bool {
	var size = len(input)
	if size == 0 {
		return false
	}

	inputList := strings.Split(input, "")
	// sorts in-place
	sort.Strings(inputList)

	var current = inputList[0]

	for i := 1; i < len(inputList); i++ {
		if current == inputList[i] {
			return false
		}
		current = inputList[i]
	}

	return true
}
