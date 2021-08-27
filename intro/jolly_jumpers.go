package intro

import "math"

const notJolly string = "Not jolly"
const jolly string = "Jolly"

func isJollyJumpers(n int, input []int) string {
	var buf = make([]int, n)

	for i := 0; i < n-1; i++ {
		var diff = int(math.Abs(float64(input[i] - input[i+1])))

		if diff < 1 || diff > n-1 || buf[diff-1] != 0 {
			return notJolly
		}

		buf[diff-1] = diff
	}

	return jolly
}
