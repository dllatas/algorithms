package main

func division(dividend int, divisor int) (int, int) {
	if divisor > dividend {
		return 0, dividend
	}
	if divisor == 0 {
		return 0, 0
	}
	var quotient = 0
	var remainder = 0
	for {
		remainder = dividend - divisor
		quotient = quotient + 1
		if remainder < divisor {
			break
		}
		dividend = remainder
	}
	return quotient, remainder
}
