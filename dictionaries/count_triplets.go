package dictionaries

func countTriplets(input []int, ratio int) int {
	var size = len(input)

	// always count 1 for each current
	var accum = make(map[int]int)

	// only store if exists in accum
	var condAccum = make(map[int]int)

	// add one to counter if exists in condAccum
	var counter = 0

	for i := size - 1; i >= 0; i-- {

		var current = input[i]
		var next = current * ratio

		if _, found := condAccum[next]; found {
			counter = counter + condAccum[next]
		}

		if _, found := accum[next]; found {
			condAccum[current] = condAccum[current] + accum[next]
		}

		accum[current] = accum[current] + 1
	}

	return counter
}
