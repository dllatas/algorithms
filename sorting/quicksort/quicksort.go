package main

import (
	"log"
)

func choosePivot(method string, low int, high int) int {
	var methodMap = map[string]int{
		"high":   high,
		"low":    low,
		"medium": (high + low) / 2,
	}
	var pivot = methodMap[method]
	return pivot
}

func partition(input []int, low int, high int, method string) int {
	var pivot = choosePivot(method, low, high)
	// first element higher than pivot
	var higher = low

	for i := low; i < high; i++ {
		if input[i] < input[pivot] {
			input[i], input[higher] = input[higher], input[i]
			log.Printf("Input iter: %v \n", input)
			higher++
		}
	}

	input[pivot], input[higher] = input[higher], input[pivot]
	log.Printf("Input part: %v \n", input)
	return higher
}

func quickSort(input []int, low int, high int, method string) []int {
	log.Printf("Input init: %v \n", input)
	if low < high {
		var pivot = partition(input, low, high, method)
		quickSort(input, low, pivot-1, method)
		quickSort(input, pivot+1, high, method)
	}
	return input
}
