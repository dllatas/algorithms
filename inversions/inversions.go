package main

type inversion struct {
	sorted []int
	count  int
}

func split(input []int) ([]int, []int) {
	var length = len(input)
	var pivot = length / 2

	var left = input[:pivot]
	var right = input[pivot:]

	return left, right
}

// sort and merge happen at the same time
func count(left []int, right []int) inversion {
	var sorted []int
	var count = 0

	for !(len(left) == 0 || len(right) == 0) {
		if left[0] <= right[0] {
			sorted = append(sorted, left[0])
			left = left[1:]
		} else {
			sorted = append(sorted, right[0])
			right = right[1:]
			count = count + len(left)
		}
	}

	// After removing the value from left or right, the for loop
	// breaks because array's lenght is 0. Therefore, we need to append
	// the remaining values to it
	if len(left) != 0 {
		sorted = append(sorted, left...)
	}

	if len(right) != 0 {
		sorted = append(sorted, right...)
	}

	return inversion{
		sorted: sorted,
		count:  count,
	}
}

// sort and count
func inversions(input []int) inversion {
	var length = len(input)

	// The recursion goes all the way down
	// until there is only one element in the input
	// At that point, it returns. This only happens at a single layer
	if length <= 1 {
		return inversion{
			sorted: input,
			count:  0,
		}
	}

	// From the previous comment, here the
	// array is divided in two pieces
	left, right := split(input)

	var leftInversion = inversions(left)
	var rightInversion = inversions(right)
	var merged = count(leftInversion.sorted, rightInversion.sorted)

	var finalCount = leftInversion.count + rightInversion.count + merged.count

	return inversion{
		sorted: merged.sorted,
		count:  finalCount,
	}
}
