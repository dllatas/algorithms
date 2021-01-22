package stack

import (
	"testing"
)

var constantMinTests = []struct {
	input []int
	want  []int
	empty bool
}{
	{
		input: []int{8, 5, 2, 9, 0, 1},
		want:  []int{0, 2, 5, 8},
		empty: false,
	},
	{
		input: []int{},
		want:  []int{},
		empty: true,
	},
	{
		input: []int{6},
		want:  []int{6},
		empty: false,
	},
	{
		input: []int{2, 4, 6, 8, 6, 4, 2},
		want:  []int{2, 2},
		empty: false,
	},
}

func TestConstantMin(t *testing.T) {
	for _, test := range constantMinTests {
		var stack = newStackMin()
		for _, input := range test.input {
			stack.push(input)
		}

		min, isEmpty := stack.getMin()
		if isEmpty != test.empty {
			t.Errorf("got: %t, want: %t", isEmpty, test.empty)
		}

		if isEmpty {
			continue
		}

		if min.value != test.want[0] {
			t.Errorf("got: %d, want: %d", min.value, test.want[0])
		}
	}
}
