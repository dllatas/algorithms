package main

import (
	"testing"
)

func TestKaratsuba(t *testing.T) {
	var tests = []struct {
		multiplicand string
		multiplier   string
		want         float64
	}{
		{
			multiplicand: "2000",
			multiplier:   "2000",
			want:         4000000.0,
		},
		{
			multiplicand: "1234",
			multiplier:   "4321",
			want:         5332114.0,
		},
		{
			multiplicand: "1234",
			multiplier:   "5678",
			want:         7006652.0,
		},
		{
			multiplicand: "3141592653589793",
			multiplier:   "2718281828459045",
			want:         8539734222673565763860525219840.000000,
		},
		{
			multiplicand: "3141592653589793238462643383279502884197169399375105820974944592",
			multiplier:   "2718281828459045235360287471352662497757247093699959574966967627",
			want:         8.539734222673568e+126,
		},
	}

	for _, tt := range tests {
		got := karatsuba(tt.multiplicand, tt.multiplier)
		if got != tt.want {
			t.Errorf("got %f, want %f", got, tt.want)
		}
	}
}
