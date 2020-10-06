package linked_list

import (
	"log"
)

type List struct {
	Value int
	Next  *List
}

func NewList(v int) *List {
	var list = &List{
		Value: v,
		Next:  nil,
	}
	return list
}

// insert
func Insert(list *List, v int) *List {
	var newList = NewList(v)

	newList.Next = list
	list = newList

	return list
}

// search
func Search(list *List, v int) *List {
	if list.Value == v {
		return list
	}
	if list.Next != nil {
		return Search(list.Next, v)
	}
	return nil
}

// deletion
func Delete(list *List, v int) {
	var listToDelete = Search(list, v)
	if listToDelete != nil {
		var prev = GetPrevious(list, v)
		if prev == nil {
			*list = *listToDelete.Next
		} else {
			prev.Next = listToDelete.Next
		}
	}
}

// prev
func GetPrevious(list *List, v int) *List {
	if list.Next == nil {
		return nil
	}
	if list.Next.Value == v {
		return list
	}
	return GetPrevious(list.Next, v)
}

// print
func Print(list *List) {
	log.Println(list.Value)
	if list.Next == nil {
		return
	}
	Print(list.Next)
}
