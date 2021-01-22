package stack

import (
	"testing"
)

var balancedTests = []struct {
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

func TestBalance(t *testing.T) {
	for _, test := range balancedTests {
		got := balance(test.input)

		if !(got == test.want) {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}

func TestBalanceUsingStack(t *testing.T) {
	for _, test := range balancedTests {
		isGood, index := balanceUsingStack(test.input)

		if !(isGood == test.want && index == test.bad) {
			t.Errorf("got: %v %d, want: %v %d", isGood, index, test.want, test.bad)
		}
	}
}
