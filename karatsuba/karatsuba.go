package main

import (
	"math"
	"strconv"
)

func rpad(s string, pad string, plength int) string {
	for i := len(s); i < plength; i++ {
		s = s + pad
	}
	return s
}

func split(operand string, point int) (string, string) {
	var highString = operand[0:point]
	var lowString = operand[point:]

	var highLen = len(highString)
	var lowLen = len(lowString)

	if highLen != lowLen {
		var size = highLen
		if highLen < lowLen {
			size = lowLen
		}
		highString = rpad(highString, "0", size)
		lowString = rpad(lowString, "0", size)
	}

	return highString, lowString
}

func sumOperands(multiplicand string, multiplier string) string {
	var a, _ = strconv.ParseFloat(multiplicand, 64)
	var b, _ = strconv.ParseFloat(multiplier, 64)
	var sum = a + b
	var sumString = strconv.FormatFloat(sum, 'f', -1, 64)
	return sumString
}

func karatsuba(multiplicand string, multiplier string) float64 {
	// Check that size of operands is lower than two
	var lowerSize = 2

	// Calculate length of operands
	var multiplicandSize = len(multiplicand)
	var multiplierSize = len(multiplier)

	if multiplicandSize < lowerSize || multiplierSize < lowerSize {
		var a, _ = strconv.ParseFloat(multiplicand, 64)
		var b, _ = strconv.ParseFloat(multiplier, 64)
		var product = a * b
		return product
	}

	// Choose the lowest one
	var minSize = multiplicandSize
	if multiplicandSize > multiplierSize {
		minSize = multiplierSize
	}

	// Find split point
	var splitPoint = int(math.Floor(float64(minSize) / 2.0))

	// Split string at split point
	highMultiplicand, lowMultiplicand := split(multiplicand, splitPoint)
	highMultiplier, lowMultiplier := split(multiplier, splitPoint)

	// Sum multiplicand and multiplier to calculate v1
	var sumMultiplicand = sumOperands(highMultiplicand, lowMultiplicand)
	var sumMultiplier = sumOperands(highMultiplier, lowMultiplier)

	// Recursive calls to karatsuba
	var z0 = karatsuba(lowMultiplicand, lowMultiplier)
	var v1 = karatsuba(sumMultiplicand, sumMultiplier)
	var v2 = karatsuba(highMultiplicand, highMultiplier)

	// calculate z1: v1 - v2 - z0
	var z1 = ((v1 - v2) - z0) * math.Pow10(splitPoint)

	// calculate z2
	var z2 = v2 * math.Pow10(splitPoint*2)

	var result = (z2 + z1) + z0

	return result
}
