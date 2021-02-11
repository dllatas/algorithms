package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
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
		var highGot = quickSort(tt.input, 0, len(tt.input)-1, "high")
		if !(reflect.DeepEqual(highGot, tt.want)) {
			t.Errorf("highGot %v, want %v", highGot, tt.want)
		}
		var lowGot = quickSort(tt.input, 0, len(tt.input)-1, "low")
		if !(reflect.DeepEqual(lowGot, tt.want)) {
			t.Errorf("lowGot %v, want %v", lowGot, tt.want)
		}

		var mediumGot = quickSort(tt.input, 0, len(tt.input)-1, "medium")
		if !(reflect.DeepEqual(mediumGot, tt.want)) {
			t.Errorf("mediumGot %v, want %v", mediumGot, tt.want)
		}
	}
}
