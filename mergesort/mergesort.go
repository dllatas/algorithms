package main

func split(input []int) ([]int, []int) {
	var length = len(input)
	var pivot = length / 2

	var left = input[:pivot]
	var right = input[pivot:]

	return left, right
}

// sort and merge happen at the same time
func merge(left []int, right []int) []int {
	var sorted []int

	for !(len(left) == 0 || len(right) == 0) {
		if left[0] <= right[0] {
			sorted = append(sorted, left[0])
			// remove element from left
			left = left[1:]
		} else {
			sorted = append(sorted, right[0])
			// remove element from right
			right = right[1:]
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

	return sorted
}

func mergeSort(input []int) []int {
	var length = len(input)

	// The recursion goes all the way down
	// until there is only one element in the input
	// At that point, it returns. This only happens at a single layer
	if length <= 1 {
		return input
	}

	// From the previous comment, here the
	// array is divided in two pieces
	left, right := split(input)

	// Then it calls the recursion with each side (left and right)
	// until it returns an array of length equals to 1
	left = mergeSort(left)
	right = mergeSort(right)

	// The merge happens at (log2(len(input)) - 1) layers
	var sorted = merge(left, right)
	return sorted
}
