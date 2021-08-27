package intro

import "testing"

func TestJollyJumpers(t *testing.T) {
	var tests = []struct {
		input []int
		want  string
	}{
		{
			input: []int{1, 4, 2, 3},
			want:  jolly,
		},
		{
			input: []int{1, 4, 2, -1, 6},
			want:  notJolly,
		},
		{
			input: []int{11, 7, 4, 2, 1, 6},
			want:  jolly,
		},
		{
			input: []int{1, 4, 2, 3, 6},
			want:  notJolly,
		},
		{
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:  notJolly,
		},
		{
			input: []int{2, 1},
			want:  jolly,
		},
		{
			input: []int{0, 4, 2, 3},
			want:  notJolly,
		},
		{
			input: []int{1, 3, 2, 4},
			want:  notJolly,
		},
		{
			input: []int{1, 4, 3, 7, 5, 10},
			want:  jolly,
		},
		{
			input: []int{2},
			want:  jolly,
		},
		{
			input: []int{1, 2},
			want:  jolly,
		},
		{
			input: []int{1},
			want:  jolly,
		},
		{
			input: []int{1, 2, 4, 7, 11, 16, 22, 29, 37, 46},
			want:  jolly,
		},
		{
			input: []int{-1, -1, -4, -7, -11, -16, -22, -29, -37, -46},
			want:  notJolly,
		},
		{
			input: []int{-1, -2, -4, -7, -11, -16, -22, -29, -37, -46},
			want:  jolly,
		},
	}

	for i, tt := range tests {
		var got = isJollyJumpers(len(tt.input), tt.input)

		if got != tt.want {
			t.Errorf("(%v): got %v, want %v", i, got, tt.want)
		}
	}
}
