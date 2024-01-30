package list

func arrayManipulation(n int32, queries [][]int32) int64 {
	var result int64 = 0
	var buffer = make([]int64, n)

	for _, query := range queries {
		a, b, k := query[0], query[1], int64(query[2])

		buffer[a-1] += k

		if b < n {
			buffer[b] += -k
		}
	}

	var accum int64 = 0
	for _, v := range buffer {
		accum = accum + v
		if accum > result {
			result = accum
		}
	}

	return result
}
