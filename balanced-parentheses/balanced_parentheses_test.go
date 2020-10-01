package balanced_parentheses

import (
	"testing"
)

var tests = []struct {
	input string
	want  bool
	bad   int
}{
	{
		input: "()",
		want:  true,
		bad:   -1,
	},
	{
		input: "(())",
		want:  true,
		bad:   -1,
	},
	{
		input: "())(",
		want:  false,
		bad:   2,
	},
	{
		input: "((())())()",
		want:  true,
		bad:   -1,
	},
	{
		input: ")()(",
		want:  false,
		bad:   0,
	},
	{
		input: "())",
		want:  false,
		bad:   2,
	},
	{
		input: "()()()()())(",
		want:  false,
		bad:   10,
	},
}

func TestNaive(t *testing.T) {
	for _, tt := range tests {
		got := naive(tt.input)

		if !(got == tt.want) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}

func TestStack(t *testing.T) {
	for _, tt := range tests {
		isGood, index := checkWithStack(tt.input)

		if !(isGood == tt.want && index == tt.bad) {
			t.Errorf("got: %v %d, want: %v %d", isGood, index, tt.want, tt.bad)
		}
	}
}
