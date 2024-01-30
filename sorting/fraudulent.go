package sorting

import (
	"fmt"
	"sort"
)

func detectFraudulent() {
	//expenditure := []int32{2, 3, 4, 2, 3, 6, 8, 4, 5}
	expenditure := []int32{10, 20, 30, 40, 50}

	var d int32 = 3
	var nots int32 = 0
	var start = 0
	var end = int(d)
	var size = len(expenditure)
	var index = d / 2
	var iter = size - int(d)
	var slice = make([]int32, d)
	var uo = make([]int32, d)
	var current = 0

	for {
		if current >= iter {
			break
		}
		fmt.Println("==================")
		fmt.Println(expenditure)

		if start == 0 {
			copy(uo, expenditure[start:end])
			copy(slice, expenditure[start:end])
			sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
		} else {
			var remove = uo[0]
			copy(uo, expenditure[start:end])

			var indexToRemove = sort.Search(int(d)-1, func(i int) bool {
				return slice[i] >= remove
			})

			slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)

			var indexToInsert = sort.Search(int(d)-1, func(i int) bool {
				return slice[i] > expenditure[end-1]
			})

			slice = append(slice, expenditure[end-1])
			copy(slice[indexToInsert+1:], slice[indexToInsert:])
			slice[indexToInsert] = expenditure[end-1]
		}

		fmt.Println(slice)
		var median int32 = 0
		if d%2 == 0 {
			median = slice[index-1] + slice[index]
		} else {
			median = slice[index] * 2
		}

		fmt.Println(expenditure[end])
		fmt.Println(median)
		if expenditure[end] >= median {
			nots += 1
		}

		start += 1
		end += 1
		current += 1
	}

	//	return nots
	fmt.Println(nots)
}
