package strings

import (
	//	"log"
	"testing"
)

type PermTest struct {
	S1   string
	S2   string
	Want bool
}

var permTests = []PermTest{
	{
		S1:   "",
		S2:   "",
		Want: true,
	},
	{
		S1:   "abc",
		S2:   "cba",
		Want: true,
	},
	{
		S1:   "abc",
		S2:   "cb",
		Want: false,
	},
	{
		S1:   "abcd",
		S2:   "dabc",
		Want: true,
	},
	{
		S1:   "test",
		S2:   "ttew",
		Want: false,
	},
	{
		S1:   "geeksforgeeks",
		S2:   "forgeeksgeeks",
		Want: true,
	},
	{
		S1:   "God",
		S2:   "dog",
		Want: false,
	},
	{
		S1:   "god ",
		S2:   "god",
		Want: false,
	},
}

func TestPermUsingHash(t *testing.T) {
	for i, test := range permTests {
		var got = permUsingHash(test.S1, test.S2)
		if got != test.Want {
			t.Errorf("test (%d): got %v, want %v", i, got, test.Want)
		}
	}
}
