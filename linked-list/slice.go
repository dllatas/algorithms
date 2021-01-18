package linked_list

import (
//	"log"
)

// seed = 0, returns last element
// seed = 1, returns before last element

func GetBeforeLast(list *List, seed int) *List {
	var size = 0
	var current = list

	for {
		size = size + 1
		if current.Next == nil {
			break
		}
		current = current.Next
	}

	var position = size - seed
	if position == 0 {
		return &List{
			Next:  nil,
			Value: 0,
		}
	}

	var index = 1
	var found = list

	for {
		if index == position {
			break
		}

		index = index + 1
		found = found.Next
	}

	return found
}
