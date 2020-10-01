package fibonnaci

import (
	"log"
	"testing"
	"time"
)

type Test []struct {
	input int
	want  int
}

var tests = Test{
	{
		input: 0,
		want:  0,
	},
	{
		input: 1,
		want:  1,
	},
	{
		input: 5,
		want:  5,
	},
	{
		input: 10,
		want:  55,
	},
	{
		input: 13,
		want:  233,
	},
	{
		input: 45,
		want:  1134903170,
	},
}

func TestRecursive(t *testing.T) {
	start := time.Now()
	for _, tt := range tests {
		got := recursive(tt.input)
		if got != tt.want {
			t.Errorf("got %v, want %vf", got, tt.want)
		}
	}
	duration := time.Since(start)
	log.Printf("Recursive duration: %s", duration)
}

func TestMemoization(t *testing.T) {
	start := time.Now()
	for _, tt := range tests {
		got := memoizationMain(tt.input, false)
		if got != tt.want {
			t.Errorf("got %v, want %v\n", got, tt.want)
		}
	}
	duration := time.Since(start)
	log.Printf("Memoization duration: %s", duration)
}

func TestDynamic(t *testing.T) {
	start := time.Now()
	for _, tt := range tests {
		got := dynamic(tt.input)
		if got != tt.want {
			t.Errorf("got %v, want %v\n", got, tt.want)
		}
	}
	duration := time.Since(start)
	log.Printf("Dynamic duration: %s", duration)
}
