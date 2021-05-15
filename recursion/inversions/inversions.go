package inversion

type inversion struct {
	sorted []int32
	count  int64
}

func split(input []int32, length int64) ([]int32, int64, []int32, int64) {
	var pivot int64 = length / int64(2)

	var left = input[:pivot]
	var right = input[pivot:]

	return left, pivot, right, length - pivot
}

func count(left []int32, leftLen int64, right []int32, rightLen int64) inversion {
	var sorted []int32
	var count int64 = 0

	for !(leftLen == 0 || rightLen == 0) {
		if left[0] <= right[0] {
			sorted = append(sorted, left[0])
			left = left[1:]
			leftLen = leftLen - 1
		} else {
			sorted = append(sorted, right[0])
			right = right[1:]
			count = count + leftLen
			rightLen = rightLen - 1
		}
	}

	// After removing the value from left or right, the for loop
	// breaks because array's lenght is 0. Therefore, we need to append
	// the remaining values to it
	if leftLen != 0 {
		sorted = append(sorted, left...)
	}

	if rightLen != 0 {
		sorted = append(sorted, right...)
	}

	return inversion{
		sorted: sorted,
		count:  count,
	}
}

func inversions(input []int32, length int64) inversion {

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
	left, leftLen, right, rightLen := split(input, length)

	var leftInversion = inversions(left, leftLen)
	var rightInversion = inversions(right, rightLen)

	var merged = count(leftInversion.sorted, leftLen, rightInversion.sorted, rightLen)

	var finalCount = leftInversion.count + rightInversion.count + merged.count

	return inversion{
		sorted: merged.sorted,
		count:  finalCount,
	}
}
