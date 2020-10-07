package linked_list

import (
//"log"
)

func RemoveDuplicates(list *List, buffer map[int]bool) {
	RemoveDuplicate(list, list, buffer)
}

func RemoveDuplicate(root *List, current *List, buffer map[int]bool) {
	_, found := buffer[current.Value]
	if found {
		Delete(root, current.Value)
	} else {
		buffer[current.Value] = true
	}

	if current.Next == nil {
		return
	}

	RemoveDuplicate(root, current.Next, buffer)
}
