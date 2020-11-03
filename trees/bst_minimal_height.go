package trees

import (
//"log"
)

func simpleMinHeightBST(bst **BinarySearchTree, input []int, size int) {
	if size == 0 {
		return
	}

	// Split at median point
	var index = size / 2
	var value = input[index]

	if *bst == nil {
		*bst = NewBinarySearchTree(value)
	} else {
		(*bst).Insert(value)
	}

	if size <= 1 {
		return
	}

	// These lines reduce some useless traverse
	var rightIndex = index + 1
	if index+1 >= size {
		rightIndex = index
	}

	left, right := input[0:index], input[rightIndex:]
	var leftSize = index
	var rightSize = size - rightIndex

	// This lines reduces a useless traverse
	if !(leftSize == 1 && left[0] == value) {
		simpleMinHeightBST(bst, left, leftSize)

	}

	// This lines reduces a useless traverse
	if !(rightSize == 1 && right[0] == value) {
		simpleMinHeightBST(bst, right, rightSize)
	}
}

// How to avoid the traverse?
func directMinHeightBST(parent, current **BinarySearchTree, input []int, size int) {
	if size == 0 {
		return
	}

	var index = size / 2
	var value = input[index]

	if *parent == nil {
		*parent = NewBinarySearchTree(value)
	}

	if *current == nil {
		*current = (*parent).NewNode(value)
	}

	if size <= 1 {
		return
	}

	// These lines reduce some useless traverse by avoiding duplicates
	var rightIndex = index + 1
	if index+1 >= size {
		rightIndex = index
	}

	left, right := input[0:index], input[rightIndex:]
	var leftSize = index
	var rightSize = size - rightIndex

	// This lines reduces a useless traverse by avoiding duplicates
	if !(leftSize == 1 && left[0] == value) {
		directMinHeightBST(current, &((*current).Left), left, leftSize)
	}

	// This lines reduces a useless traverse by avoiding duplicates
	if !(rightSize == 1 && right[0] == value) {
		directMinHeightBST(current, &((*current).Right), right, rightSize)
	}
}
