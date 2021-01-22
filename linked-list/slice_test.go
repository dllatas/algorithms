package linkedlist

import (
	"reflect"
	"testing"
)

type SliceToLastTest []struct {
	from  int
	input *List
	want  []int
}

func TestSliceToLast(t *testing.T) {
	var tests = SliceToLastTest{
		{
			from:  0,
			input: NewListFromSlice([]int{5, 4, 3}),
			want:  []int{},
		},
		{
			from:  1,
			input: NewListFromSlice([]int{5, 4, 6}),
			want:  []int{6},
		},
		{
			from:  2,
			input: NewListFromSlice([]int{5, 4, 6}),
			want:  []int{4, 6},
		},
		{
			from:  3,
			input: NewListFromSlice([]int{5, 4, 6}),
			want:  []int{5, 4, 6},
		},
		{
			from:  5,
			input: NewListFromSlice([]int{5, 4, 6}),
			want:  []int{5, 4, 6},
		},
	}

	for i, test := range tests {
		_, list := test.input.sliceToLast(test.from)
		var got = list.GetValues()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("(%d) got %v, want %v", i, got, test.want)
		}
	}
}

type LengthTest []struct {
	input []int
	list  *List
	want  int
}

func TestLength(t *testing.T) {
	var tests = LengthTest{
		{
			input: []int{5, 4, 6},
		},
		{
			input: []int{},
		},
		{
			input: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, test := range tests {
		test.list = NewListFromSlice(test.input)

		var length = len(test.input)
		test.want = length
		if length == 0 {
			test.want = 1
		}

		var got = test.list.length()
		if got != test.want {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}
