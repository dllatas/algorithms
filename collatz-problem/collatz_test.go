package main

import (
	"reflect"
	"testing"
)

func TestCollatz(t *testing.T) {
	var tests = []struct {
		input int
		iter  int
		seq   []int
	}{
		{
			input: 12,
			iter:  10,
			seq:   []int{12, 6, 3, 10, 5, 16, 8, 4, 2, 1},
		},
		{
			input: 19,
			iter:  21,
			seq:   []int{19, 58, 29, 88, 44, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1},
		},
		{
			input: 22,
			iter:  16,
			seq:   []int{22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1},
		},
	}

	for _, tt := range tests {
		iter, seq := collatz(tt.input)

		if iter != tt.iter {
			t.Errorf("got %v, want %v", iter, tt.iter)
		}

		if !reflect.DeepEqual(seq, tt.seq) {
			t.Errorf("got %v, want %v", seq, tt.seq)
		}
	}
}
