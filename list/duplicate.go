package list

import (
//	"log"
)

func removeDuplicates(list *List, buffer map[int]bool) {
	removeDuplicate(list, list, buffer)
}

func removeDuplicate(root *List, current *List, buffer map[int]bool) {
	_, found := buffer[current.Value]
	if found {
		root.Delete(current.Value)
	} else {
		buffer[current.Value] = true
	}

	if current.Next == nil {
		return
	}

	removeDuplicate(root, current.Next, buffer)
}

func removeDuplicateNoBuffer(list *List) {
	var current *List

	for {
		var toRemove = list.Value
		var removed = false
		current = list

		for {
			if current.Next == nil {
				break
			}

			if toRemove == current.Next.Value {
				list.Delete(toRemove)
				removed = true
			}

			if list.Next == nil {
				break
			}

			if current.Next.Next == nil {
				break
			}

			current = current.Next
		}

		if list.Next == nil {
			break
		}

		if !removed {
			list = list.Next
		}
	}
}
