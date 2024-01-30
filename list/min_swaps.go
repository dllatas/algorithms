package list

import (
	"log"
)

func minimumSwaps(arr []int) int {
	var swaps = 0
	var size = len(arr)

	for i := 0; i < size; i++ {
		for {
			if arr[i]-1 == i {
				break
			}

			log.Println(arr)
			log.Println(i)

			log.Println(arr[i])
			log.Println(arr[arr[i]-1])

			arr[arr[i]-1], arr[i] = arr[i], arr[arr[i]-1]
			swaps = swaps + 1
		}
	}

	log.Println(arr)
	return swaps
}

/*
def minimumSwaps(arr):
    swaps = 0
    n = len(arr)

    for idx in xrange(n):
        while arr[idx]-1 != idx:
            ele = arr[idx]
            arr[ele-1], arr[idx] = arr[idx], arr[ele-1]
            swaps += 1
    return swaps
*/
