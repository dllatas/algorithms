package main

func collatz(input int) (int, []int) {
	if input == 1 {
		return 1, []int{input}
	}

	var transformed = 0
	var remainder = input % 2
	if remainder == 0 {
		transformed = input / 2
	} else {
		transformed = (3 * input) + 1
	}

	iteration, serie := collatz(transformed)

	iteration = iteration + 1
	serie = append([]int{input}, serie...)

	return iteration, serie
}
