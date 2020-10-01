package karatsuba

import (
	"log"
	"math"
	"strconv"
)

// Check that size of operands is lower than two
const lowerSize = 2
const padChar = "0"

type operand struct {
	value string
	size  int
}

func newOperand(value string, size int) operand {
	if size != 0 {
		return operand{
			value: value,
			size:  size,
		}
	}
	return operand{
		value: value,
		size:  len(value),
	}
}

func rpad(s string, size int, pad string, plength int) (string, int) {
	var counter = 0
	for i := size; i < plength; i++ {
		s = s + pad
		counter = counter + 1
	}
	return s, counter
}

func split(operand operand, point int) (operand, operand) {
	var highPad = 0
	var lowPad = 0

	var highString = operand.value[:point]
	var lowString = operand.value[point:]

	var highLen = point
	var lowLen = operand.size - point

	if highLen != lowLen {
		var size = highLen
		if highLen < lowLen {
			size = lowLen
		}
		highString, highPad = rpad(highString, highLen, padChar, size)
		lowString, lowPad = rpad(lowString, lowLen, padChar, size)
	}

	var high = newOperand(highString, highLen+highPad)
	var low = newOperand(lowString, lowLen+lowPad)

	return high, low
}

func sum(operD operand, operR operand) operand {
	var result = apply([]operand{operD, operR}, "+")
	var resultString = strconv.FormatFloat(result, 'f', -1, 64)
	return operand{
		value: resultString,
		size:  len(resultString),
	}
}

func getMinSize(multiplicandSize, multiplierSize int) int {
	var minSize = multiplicandSize
	if multiplicandSize > multiplierSize {
		minSize = multiplierSize
	}
	return minSize
}

func getSplitPoint(minSize int) int {
	var minSizeFloat = float64(minSize)
	var halfMinSize = minSizeFloat / 2.0
	var floor = math.Floor(halfMinSize)
	return int(floor)
}

func apply(operands []operand, operation string) float64 {
	var result = 0.0
	for index, operand := range operands {
		var parsed, _ = strconv.ParseFloat(operand.value, 64)
		switch operation {
		case "+":
			result = result + parsed
		case "*":
			if index == 0 {
				result = parsed
				continue
			}
			result = result * parsed
		default:
			log.Printf("Operation not supported: %s\n", operation)
		}
	}
	return result
}

func karatsuba(operD operand, operR operand) float64 {
	if operD.size < lowerSize || operR.size < lowerSize {
		var product = apply([]operand{operD, operR}, "*")
		return product
	}

	// Choose the lowest one
	var minSize = getMinSize(operD.size, operR.size)

	// Find split point
	var splitPoint = getSplitPoint(minSize)

	// Split string at split point
	dHigh, dLow := split(operD, splitPoint)
	rHigh, rLow := split(operR, splitPoint)

	// Sum multiplicand and multiplier to calculate v1
	var sumD = sum(dHigh, dLow)
	var sumR = sum(rHigh, rLow)

	// Recursive calls to karatsuba
	var z0 = karatsuba(dLow, rLow)
	var v1 = karatsuba(sumD, sumR)
	var v2 = karatsuba(dHigh, rHigh)

	// calculate z1: v1 - v2 - z0
	var z1 = ((v1 - v2) - z0) * math.Pow10(splitPoint)

	var z2 = v2 * math.Pow10(splitPoint*2)

	var result = (z2 + z1) + z0
	return result
}
