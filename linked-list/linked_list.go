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

func NewListFromSlice(v []int) *List {
	if len(v) == 0 {
		return &List{Next: nil}
	}

	var list = NewList(v[0])
	for i := 1; i < len(v); i++ {
		list = Insert(list, v[i])
	}

	return list
}

func Insert(list *List, v int) *List {
	var newList = NewList(v)

	newList.Next = list
	list = newList

	return list
}

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
		var prev = GetPrevious(list, listToDelete)
		if prev == nil {
			if listToDelete.Next != nil {
				*list = *listToDelete.Next
			}
		} else {
			prev.Next = listToDelete.Next
		}
	}
}

func GetPrevious(list *List, v *List) *List {
	if list.Next == nil {
		return nil
	}

	if list.Next == v {
		return list
	}
	return GetPrevious(list.Next, v)
}

func GetValues(list *List) []int {
	var value []int
	var current = list
	for {
		value = append(value, current.Value)
		if current.Next == nil {
			break
		}
		current = current.Next
	}
	return value
}

func Print(list *List) {
	var value []List
	var current = list
	for {
		value = append(value, *current)
		if current.Next == nil {
			break
		}
		current = current.Next
	}

	log.Printf("List: %v", value)
}
