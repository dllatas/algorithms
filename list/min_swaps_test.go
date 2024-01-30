package list

import (
	//"log"
	"reflect"
	"testing"
)

type MinSwapTests []struct {
	input []int
	//want  []int
	want int
}

var minSwapTests = MinSwapTests{
	{
		input: []int{4, 3, 1, 2},
		want:  3,
	},
}

func TestMinSwaps(t *testing.T) {
	for _, test := range minSwapTests {

		var swaps = minimumSwaps(test.input)

		if !reflect.DeepEqual(swaps, test.want) {
			t.Errorf("got %v, want %v", swaps, test.want)
		}
	}
}
