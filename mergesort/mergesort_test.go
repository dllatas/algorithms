package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{
			input: []int{},
			want:  []int{},
		},
		{
			input: []int{9},
			want:  []int{9},
		},
		{
			input: []int{9, 4, 6},
			want:  []int{4, 6, 9},
		},
		{
			input: []int{9, 4, 2, 6},
			want:  []int{2, 4, 6, 9},
		},
		{
			input: []int{2, 4, 6, 9},
			want:  []int{2, 4, 6, 9},
		},
		{
			input: []int{2, 2, 6, 6},
			want:  []int{2, 2, 6, 6},
		},
		{
			input: []int{38, 27, 43, 3, 9, 82, 10},
			want:  []int{3, 9, 10, 27, 38, 43, 82},
		},
		{
			input: []int{12, 11, 13, 5, 6, 7},
			want:  []int{5, 6, 7, 11, 12, 13},
		},
	}

	for _, tt := range tests {
		got := mergeSort(tt.input)

		if !(reflect.DeepEqual(got, tt.want)) {
			t.Errorf("got %v, want %vf", got, tt.want)
		}
	}
}
