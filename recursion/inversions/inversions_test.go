package inversion

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{
			input: []int{9},
			want:  0,
		},
		{
			input: []int{9, 4, 6},
			want:  2,
		},
		{
			input: []int{9, 4, 2, 6},
			want:  4,
		},
		{
			input: []int{2, 4, 6, 9},
			want:  0,
		},
	}

	for _, tt := range tests {
		got := inversions(tt.input, len(tt.input))

		if !(reflect.DeepEqual(got.count, tt.want)) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
