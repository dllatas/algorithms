package fibonnaci

// bottom-up
func dynamic(input int) int {
	if input == 0 {
		return input
	}

	var prevIter = 1
	var beforePrevIter = 0

	for iter := 2; iter < input; iter++ {
		var currentIter = prevIter + beforePrevIter
		beforePrevIter = prevIter
		prevIter = currentIter
	}

	return prevIter + beforePrevIter
}
