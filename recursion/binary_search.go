package recursion

import "log"

func binarySearch(nums []int, target int) int {
	var size = len(nums)
	var index = size / 2
	var prevIndex = 0

	var buff = make(map[int]bool)

	for {
		log.Printf("%d - %d", index, nums[index])

		_, ok := buff[index]
		if ok {
			return -1
		}

		buff[index] = true
		if nums[index] == target {
			return index
		}

		if nums[index] < target {
			prevIndex = index
			index = (size + index) / 2
			continue
		}

		if nums[index] > target {
			size = index
			index = (index + prevIndex) / 2
			continue
		}
	}
}

func classicBinarySearch(nums []int, target int) int {
	var left = 0
	var right = len(nums) - 1

	for left <= right {
		var middle = (left + right) / 2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			right = middle - 1
		}

		if nums[middle] < target {
			left = middle + 1
		}
	}

	return -1
}
