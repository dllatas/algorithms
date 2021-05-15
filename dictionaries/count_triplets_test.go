package dictionaries

import (
	"testing"
)

type countTripletsFixture struct {
	input []int
	ratio int
	want  int
}

func newCountTripletsFixture(input []int, ratio, want int) countTripletsFixture {
	return countTripletsFixture{
		input: input,
		ratio: ratio,
		want:  want,
	}
}

func TestCountTriplets(t *testing.T) {
	var tests []countTripletsFixture

	var testA = newCountTripletsFixture([]int{3, 9, 27, 81}, 3, 2)
	tests = append(tests, testA)

	var testB = newCountTripletsFixture([]int{2, 4, 8, 50}, 2, 1)
	tests = append(tests, testB)

	var testC = newCountTripletsFixture([]int{1, 4, 16, 64}, 4, 2)
	tests = append(tests, testC)

	var testD = newCountTripletsFixture([]int{1, 2, 2, 4}, 2, 2)
	tests = append(tests, testD)

	var testE = newCountTripletsFixture([]int{1, 3, 9, 9, 27, 81}, 3, 6)
	tests = append(tests, testE)

	var testF = newCountTripletsFixture([]int{1, 5, 5, 25, 125}, 5, 4)
	tests = append(tests, testF)

	for _, tt := range tests {
		got := countTriplets(tt.input, tt.ratio)
		if got != tt.want {
			t.Errorf("got %v, want %vf", got, tt.want)
		}
	}
}
