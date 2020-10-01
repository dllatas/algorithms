package permutation

import (
	"reflect"
	"sort"
	"strings"
)

func sortStringIntoArray(input string) []string {
	inputList := strings.Split(input, "")
	sort.Strings(inputList)
	return inputList
}

func isStringPermutation(input string, sec string) bool {
	var inputList = sortStringIntoArray(input)
	var secList = sortStringIntoArray(sec)

	if reflect.DeepEqual(inputList, secList) {
		return true
	}

	return false
}
