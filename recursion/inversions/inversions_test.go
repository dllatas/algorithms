package inversion

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var tests = []struct {
		input []int32
		want  int64
	}{
		{
			input: []int32{9},
			want:  0,
		},
		{
			input: []int32{9, 4, 6},
			want:  2,
		},
		{
			input: []int32{9, 4, 2, 6},
			want:  4,
		},
		{
			input: []int32{2, 4, 6, 9},
			want:  0,
		},
	}

	for _, tt := range tests {
		got := inversions(tt.input, int64(len(tt.input)))

		if !(reflect.DeepEqual(got.count, tt.want)) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
