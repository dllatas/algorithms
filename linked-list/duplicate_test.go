package linked_list

import (
	//"log"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	type Tests []struct {
		input []int
		want  []int
	}
	var tests = Tests{
		{
			input: []int{1, 1},
			want:  []int{1},
		},
		{
			input: []int{1, 1, 2, 2},
			want:  []int{2, 1},
		},
		{
			input: []int{1, 2, 3, 4, 5, 6, 4},
			want:  []int{6, 5, 4, 3, 2, 1},
		},
	}

	for _, test := range tests {
		var buffer = make(map[int]bool)
		var list = NewListFromSlice(test.input)

		RemoveDuplicates(list, buffer)
		var got = GetValues(list)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}
