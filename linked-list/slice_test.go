package linked_list

import (
	"reflect"
	"testing"
)

type SliceTest []struct {
	seed  int
	input []int
	want  int
}

func TestSlice(t *testing.T) {
	var tests = SliceTest{
		{
			seed:  0,
			input: []int{5, 4, 3},
			want:  5,
		},
		{
			seed:  1,
			input: []int{5, 4, 6},
			want:  4,
		},
		{
			seed:  2,
			input: []int{5, 4, 6},
			want:  6,
		},
		{
			seed:  3,
			input: []int{5, 4, 6},
			want:  0,
		},
	}

	for _, test := range tests {
		var list = NewListFromSlice(test.input)
		var sliced = GetBeforeLast(list, test.seed)
		if !reflect.DeepEqual(sliced.Value, test.want) {
			t.Errorf("got %v, want %v", sliced.Value, test.want)
		}
	}
}
