func climbStairs(n int) int {
	var steps = []int{1, 2}
	var counter = 0

	for _, step := range steps {
		counter += count(step, n, steps)
	}

	return counter
}

func count(step, n int, steps []int) int {
	var diff = n - step

	if diff < 0 {
		return 0
	}

	if diff == 0 {
		return 1
	}

	var counter = 0
	for _, step := range steps {
		counter += count(step, diff, steps)
	}

	return counter
}
