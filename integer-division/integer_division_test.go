package main

import (
	"testing"
)

func TestDivision(t *testing.T) {
	var tests = []struct {
		dividend  int
		divisor   int
		quotient  int
		remainder int
	}{
		{
			dividend:  20,
			divisor:   5,
			quotient:  4,
			remainder: 0,
		},
		{
			dividend:  3,
			divisor:   6,
			quotient:  0,
			remainder: 3,
		},
		{
			dividend:  7,
			divisor:   6,
			quotient:  1,
			remainder: 1,
		},
		{
			dividend:  0,
			divisor:   0,
			quotient:  0,
			remainder: 0,
		},
		{
			dividend:  0,
			divisor:   8,
			quotient:  0,
			remainder: 0,
		},
		{
			dividend:  100000000,
			divisor:   10,
			quotient:  10000000,
			remainder: 0,
		},
	}

	for _, tt := range tests {
		gotQuotient, gotRemainder := division(tt.dividend, tt.divisor)

		if gotQuotient != tt.quotient {
			t.Errorf("got %v, want %v", gotQuotient, tt.quotient)
		}

		if gotRemainder != tt.remainder {
			t.Errorf("got %v, want %v", gotRemainder, tt.remainder)
		}
	}
}
