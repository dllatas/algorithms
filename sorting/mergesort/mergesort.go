package mergesort

import (
	"log"
)

func split(input []int, length int) ([]int, int, []int, int) {
	var pivot = length / 2

	var left = input[:pivot]
	var right = input[pivot:]

	return left, pivot, right, length - pivot
}

// sort and merge happen at the same time
func merge(left []int, leftLength int, right []int, rightLength int, debug bool) []int {
	var sorted []int

	// While the lengths for both sides are diff than 0
	for !(leftLength == 0 || rightLength == 0) {
		if debug {
			log.Printf("merge left: %v -> %d\n", left, leftLength)
			log.Printf("merge right: %v -> %d\n", right, rightLength)
		}

		// move first element from left side to sorted
		// else move first element from right side to sorted
		if left[0] <= right[0] {
			sorted = append(sorted, left[0])
			left = left[1:]
			leftLength = leftLength - 1

		} else {
			sorted = append(sorted, right[0])
			right = right[1:]
			rightLength = rightLength - 1
		}
	}

	// Eventually the previous loop will break
	// after removing values from either left or right or both
	// since at least one of those sides will have length equal to 0.
	// Therefore, the other side still has some values that we just
	// append to the sorted array
	if leftLength != 0 {
		sorted = append(sorted, left...)
	}

	if rightLength != 0 {
		sorted = append(sorted, right...)
	}

	if debug {
		log.Printf("sorted: %v -> %d\n", sorted, len(sorted))
	}

	return sorted
}

func mergeSort(input []int, length int, debug bool) []int {
	if debug {
		log.Printf("%v -> %d\n", input, length)
	}

	// The recursion goes all the way down
	// until there is only one element in the input
	// At that point, it returns. This only happens at a single layer
	if length <= 1 {
		return input
	}

	// From the previous comment, here the
	// array is divided in two pieces
	left, leftLength, right, rightLength := split(input, length)

	// Then it calls the recursion with each side (left and right)
	// until it returns an array of length equals to 1
	left = mergeSort(left, leftLength, debug)
	right = mergeSort(right, rightLength, debug)

	if debug {
		log.Printf("split: %v -> %d\n", left, leftLength)
		log.Printf("split: %v -> %d\n", right, rightLength)
	}

	// The merge happens at (log2(len(input)) - 1) layers
	var sorted = merge(left, leftLength, right, rightLength, debug)
	return sorted
}
