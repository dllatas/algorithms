package linkedlist

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
	var size = len(v)
	if size == 0 {
		return &List{Next: nil}
	}

	var list = NewList(v[size-1])
	for i := size - 2; i >= 0; i-- {
		list = Insert(list, v[i])
	}

	return list
}

// Insert -> adds at the beginning
func Insert(list *List, v int) *List {
	var newList = NewList(v)

	newList.Next = list
	list = newList

	return list
}

func (list *List) Search(v int) *List {
	if list.Value == v {
		return list
	}
	if list.Next != nil {
		return list.Next.Search(v)
	}
	return nil
}

func (list *List) Delete(v int) {
	var listToDelete = list.Search(v)
	if listToDelete != nil {
		var prev = list.GetPrevious(listToDelete)
		if prev == nil {
			if listToDelete.Next != nil {
				*list = *listToDelete.Next
			}
		} else {
			prev.Next = listToDelete.Next
		}
	}
}

func (list *List) GetPrevious(v *List) *List {
	if list.Next == nil {
		return nil
	}

	if list.Next == v {
		return list
	}
	return list.Next.GetPrevious(v)
}

func (list *List) GetValues() []int {
	var value = make([]int, 0)
	if list.Value == 0 && list.Next == nil {
		return value
	}

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

func (list *List) Print(mode string) {
	var value []List
	var valueRaw []*List
	var current = list
	for {
		if mode == "raw" {
			valueRaw = append(valueRaw, current)
		} else {
			value = append(value, *current)
		}
		if current.Next == nil {
			break
		}
		current = current.Next
	}

	if mode == "raw" {
		log.Printf("List Raw %v", valueRaw)
	} else {
		log.Printf("List: %v", value)

	}
}
