package puzzle

import (
	//"log"
	"testing"
)

type GetHeavierPillTest struct {
	NumberOfBottles int
	HeaviestBottle  int
}

func TestGetHeavierPill(t *testing.T) {
	var tests = []GetHeavierPillTest{
		{NumberOfBottles: 0, HeaviestBottle: 0},
		{NumberOfBottles: 1, HeaviestBottle: 1},
		{NumberOfBottles: 2, HeaviestBottle: 1},
		{NumberOfBottles: 20, HeaviestBottle: 5},
	}

	for i, test := range tests {
		var got = getHeavierBottle(test.NumberOfBottles, test.HeaviestBottle)
		var want = test.HeaviestBottle

		if got.Pills != want {
			t.Errorf("test (%d) -> got: %v, want: %v", i, got.Pills, want)
		}
	}
}
