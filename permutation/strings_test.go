package permutation

import (
	"testing"
)

var tests = []struct {
	input string
	sec   string
	want  bool
}{
	{
		input: "aaa",
		sec:   "aaa",
		want:  true,
	},
	{
		input: "abc",
		sec:   "cba",
		want:  true,
	},
	{
		input: "",
		sec:   "",
		want:  true,
	},
	{
		input: "xyz",
		sec:   "abc",
		want:  false,
	},
	{
		input: "abcd",
		sec:   "dcbaa",
		want:  false,
	},
}

func TestUniqueSort(t *testing.T) {
	for _, tt := range tests {
		var got = isStringPermutation(tt.input, tt.sec)

		if got != tt.want {
			t.Errorf("got: %t, want: %t. For input (%s, %s)", got, tt.want, tt.input, tt.sec)
		}
	}
}
