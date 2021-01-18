package bit

import (
	"testing"
)

type AddTests struct {
	SummandA string
	SummandB string
	Want     string
}

func TestAdd(t *testing.T) {
	var tests = []AddTests{
		{
			SummandA: "",
			SummandB: "",
			Want:     "",
		},
		{
			SummandA: "0110",
			SummandB: "",
			Want:     "0110",
		},
		{
			SummandA: "0110",
			SummandB: "0010",
			Want:     "1000",
		},
		{
			SummandA: "1111",
			SummandB: "1111",
			Want:     "11110",
		},
	}

	for i, test := range tests {
		var got = Add(test.SummandA, test.SummandB)
		if got != test.Want {
			t.Errorf("test (%d) -> got: %v, want: %v", i, got, test.Want)
		}
	}
}

type PadTests struct {
	Summand string
	MaxSize int
	PadChar string
	Want    string
}

func TestPad(t *testing.T) {
	var tests = []PadTests{
		{
			Summand: "",
			MaxSize: 4,
			PadChar: "0",
			Want:    "0000",
		},
		{
			Summand: "0110",
			MaxSize: 4,
			PadChar: "0",
			Want:    "0110",
		},
		{
			Summand: "0010",
			MaxSize: 4,
			PadChar: "0",
			Want:    "0010",
		},
	}
	for i, test := range tests {
		var pad = NewPad(test.Summand, "", len(test.Summand), test.MaxSize)
		pad.Fill()
		if pad.SummandToFill != test.Want {
			t.Errorf("test (%d) -> got: %v, want: %v", i, pad.SummandToFill, test.Want)
		}
	}
}

type SumTest struct {
	a, b, c   string
	WantSum   string
	WantCarry string
}

func TestSum(t *testing.T) {
	var tests = []SumTest{
		{a: "0", b: "0", c: "1", WantSum: "1", WantCarry: "0"},
		{a: "0", b: "1", c: "0", WantSum: "1", WantCarry: "0"},
		{a: "1", b: "0", c: "0", WantSum: "1", WantCarry: "0"},
		{a: "1", b: "1", c: "1", WantSum: "1", WantCarry: "1"},
		{a: "0", b: "1", c: "1", WantSum: "0", WantCarry: "1"},
		{a: "1", b: "0", c: "1", WantSum: "0", WantCarry: "1"},
		{a: "1", b: "1", c: "0", WantSum: "0", WantCarry: "1"},
	}
	for i, test := range tests {
		valueGot, carryGot := sum(test.a, test.b, test.c)
		if valueGot != test.WantSum {
			t.Errorf("test value (%d) -> got: %v, want: %v", i, valueGot, test.WantSum)
		}
		if carryGot != test.WantCarry {
			t.Errorf("test carry (%d) -> got: %v, want: %v", i, carryGot, test.WantCarry)
		}
	}
}
