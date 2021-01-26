package list

import (
	//"log"
	"reflect"
	"testing"
)

type DuplicatesTests []struct {
	input []int
	want  []int
}

var duplicatedTests = DuplicatesTests{
	{
		input: []int{1, 1},
		want:  []int{1},
	},
	{
		input: []int{1, 1, 2, 2},
		want:  []int{1, 2},
	},
	{
		input: []int{1, 2, 3, 4, 5, 6, 4},
		want:  []int{1, 2, 3, 5, 6, 4},
	},
	{
		input: []int{3, 4, 4, 3},
		want:  []int{4, 3},
	},
	{
		input: []int{5, 5, 5, 6},
		want:  []int{5, 6},
	},
	{
		input: []int{8, 7, 3, 5},
		want:  []int{8, 7, 3, 5},
	},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, test := range duplicatedTests {
		var buffer = make(map[int]bool)
		var list = NewListFromSlice(test.input)

		removeDuplicates(list, buffer)
		var got = list.GetValues()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}

func TestRemoveDuplicatesNoBuffer(t *testing.T) {
	for _, test := range duplicatedTests {
		var list = NewListFromSlice(test.input)

		removeDuplicateNoBuffer(list)
		var got = list.GetValues()

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}
