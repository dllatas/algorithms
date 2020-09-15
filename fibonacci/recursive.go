package fibonnaci

func recursive(input int) int {
	if input == 0 || input == 1 {
		return input
	}
	return recursive(input-1) + recursive(input-2)
}
