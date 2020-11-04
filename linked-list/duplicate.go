package linked_list

import (
//	"log"
)

func RemoveDuplicates(list *List, buffer map[int]bool) {
	RemoveDuplicate(list, list, buffer)
}

func RemoveDuplicate(root *List, current *List, buffer map[int]bool) {
	_, found := buffer[current.Value]
	if found {
		root.Delete(current.Value)
	} else {
		buffer[current.Value] = true
	}

	if current.Next == nil {
		return
	}

	RemoveDuplicate(root, current.Next, buffer)
}

func RemoveDuplicateNoBuffer(list *List) {
	var current *List

	for {
		//log.Println("--------new big iter----------")
		var toRemove = list.Value
		var removed = false
		current = list

		/*log.Printf("list: %v", list.GetValues())
		log.Printf("toRemove: %v", toRemove)
		log.Printf("current: %v", current.GetValues())*/

		for {

			/*log.Println("--------new small iter----------")
			log.Printf("current: %v", current.GetValues())*/

			if current.Next == nil {
				break
			}

			if toRemove == current.Next.Value {
				list.Delete(toRemove)
				removed = true
				/*log.Printf("Deleted: %d", toRemove)
				log.Printf("list: %v", list.GetValues())*/
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
