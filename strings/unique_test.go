package strings

import (
	"log"
	"testing"
)

type UniqueTest struct {
	Value string
	Want  bool
}

var uniqueTests = []UniqueTest{
	{
		Value: "",
		Want:  true,
	},
	{
		Value: "a",
		Want:  true,
	},
	{
		Value: "xx",
		Want:  false,
	},
	{
		Value: "jklzz",
		Want:  false,
	},
	{
		Value: "abcdefgh",
		Want:  true,
	},
	{
		Value: "abcabc",
		Want:  false,
	},
	{
		Value: "abcdefghijklmnopqrstuvwxyza",
		Want:  false,
	},
	{
		Value: "abcdefghijklmnopqrstuvwxyz",
		Want:  true,
	},
}

func TestUniqueUsingHash(t *testing.T) {
	for _, test := range uniqueTests {
		var got = uniqueUsingHash(test.Value)
		if got != test.Want {
			t.Errorf("got %v, want %v", got, test.Want)
		}
	}
}

func TestUnique(t *testing.T) {
	for _, test := range uniqueTests {
		var got = unique(test.Value)
		if got != test.Want {
			t.Errorf("got %v, want %v", got, test.Want)
		}
	}
}

func TestUniqueUsingBitVector(t *testing.T) {
	for _, test := range uniqueTests {
		var got = uniqueUsingBitVector(test.Value)
		if got != test.Want {
			t.Errorf("got %v, want %v", got, test.Want)
		}
	}
}

func TestUniqueUsingSort(t *testing.T) {
	for _, test := range uniqueTests {
		var got = uniqueUsingSort(test.Value)
		if got != test.Want {
			log.Println(test.Value)
			t.Errorf("got %v, want %v", got, test.Want)
		}
	}
}
