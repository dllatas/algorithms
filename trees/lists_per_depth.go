package trees

import (
	"github.com/dllatas/algorithms/list"
	"log"
)

func getLists(tree *BinarySearchTree) map[int]*list.List {
	var height = tree.Height()
	var lists = make(map[int]*list.List)

	for i := 1; i <= height; i++ {
		log.Printf("%v", lists)
		var values []int
		traverseUntil(tree, i, 1, &values)
		var list = list.NewListFromSlice(values)
		lists[i] = list
	}

	printLists(lists, height)

	return lists
}

func traverseUntil(tree *BinarySearchTree, height, current int, values *[]int) {
	if tree == nil {
		return
	}

	if current == height {
		*values = append(*values, tree.Value)
		return
	}

	current = current + 1

	traverseUntil(tree.Left, height, current, values)
	traverseUntil(tree.Right, height, current, values)
}

func printLists(lists map[int]*list.List, height int) {
	for i := 1; i <= height; i++ {
		log.Println(i)
		lists[i].Print("")
	}
}
