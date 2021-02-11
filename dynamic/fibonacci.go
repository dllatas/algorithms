package dynamic

import (
	"log"
)

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

func recursive(input int) int {
	if input == 0 || input == 1 {
		return input
	}
	return recursive(input-1) + recursive(input-2)
}

// top-down
func memoizationMain(input int, debug bool) int {
	memo := make([]int, input+1)
	return memoization(input, memo, debug)
}

func memoization(input int, memo []int, debug bool) int {
	if input == 0 || input == 1 {
		return input
	}

	if debug {
		log.Printf("pre -> input: %v, fib: %v, memo: %v", input, memo[input], memo)
	}

	if memo[input] == 0 {
		memo[input] = memoization(input-1, memo, debug) + memoization(input-2, memo, debug)
	}

	if debug {
		log.Printf("post -> input: %v, fib: %v, memo: %v", input, memo[input], memo)
	}

	return memo[input]
}
