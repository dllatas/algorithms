package bit

import (
	"log"
)

type Pad struct {
	SummandToFill string
	Summand       string
	Size          int
	DesiredSize   int
	Char          string
}

func NewPad(toFill, summand string, size, maxSize int) *Pad {
	return &Pad{
		SummandToFill: toFill,
		Size:          size,
		Summand:       summand,
		DesiredSize:   maxSize,
		Char:          "0",
	}
}

func (p *Pad) Fill() {
	var diff = p.DesiredSize - p.Size

	var word = ""
	for i := 0; i < diff; i++ {
		word = word + p.Char
	}

	log.Println(word)

	p.SummandToFill = word + p.SummandToFill
}

func Add(a, b string) string {
	sizeA, sizeB := len(a), len(b)
	if sizeA+sizeB == 0 {
		return ""
	}

	var pad = NewPad(a, b, sizeA, sizeB)
	if sizeA > sizeB {
		pad = NewPad(b, a, sizeB, sizeA)
	}
	pad.Fill()

	var result = ""
	var carry = "0"
	for i := pad.DesiredSize - 1; i >= 0; i-- {
		var value string
		x, y := pad.Summand[i], pad.SummandToFill[i]
		value, carry = sum(string(x), string(y), carry)
		result = value + result
	}

	if carry != "0" {
		result = carry + result
	}

	return result
}

func sum(a, b, c string) (string, string) {
	var bitArray = []string{a, b, c}

	var numberOfZeros = 0
	var numberOfOnes = 0

	for _, bit := range bitArray {
		if bit == "0" {
			numberOfZeros = numberOfZeros + 1
			continue
		}
		numberOfOnes = numberOfOnes + 1
	}

	var value = "0"
	if numberOfZeros == 2 || numberOfOnes == 3 {
		value = "1"
	}

	var carry = "0"
	if numberOfOnes >= 2 {
		carry = "1"
	}

	return value, carry
}
