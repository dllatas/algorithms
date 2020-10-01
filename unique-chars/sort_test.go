package unique_chars

import (
	"testing"
)

var tests = []struct {
	input string
	want  bool
}{
	{
		input: "aaa",
		want:  false,
	},
	{
		input: "abc",
		want:  true,
	},
	{
		input: "",
		want:  false,
	},
	{
		input: "abcdefghijklmnopqrstuvwxyza",
		want:  false,
	},
	{
		input: "abcdefghijklmnopqrstuvwxyz",
		want:  true,
	},
}

func TestUniqueSort(t *testing.T) {
	for _, tt := range tests {
		var got = uniqueSort(tt.input)

		if got != tt.want {
			t.Errorf("got: %t, want: %t", got, tt.want)
		}
	}
}
