package trees

import (
	//"log"
	"errors"
	//	"math"
	//	"fmt"
)

type BinarySearchTree struct {
	Value  int
	Parent *BinarySearchTree
	Left   *BinarySearchTree
	Right  *BinarySearchTree
}

func NewBinarySearchTree(value int) *BinarySearchTree {
	var binarySearchTree = &BinarySearchTree{
		Value: value,
	}
	return binarySearchTree
}

func (t *BinarySearchTree) NewNode(value int) *BinarySearchTree {
	var node = NewBinarySearchTree(value)
	node.Parent = t
	return node
}

func (t *BinarySearchTree) Insert(value int) {
	Insert(t, &t, value)
}

func Insert(parent *BinarySearchTree, current **BinarySearchTree, value int) {
	if *current == nil {
		*current = parent.NewNode(value)
		return
	}

	if value == (*current).Value {
		return
	}

	if value < (*current).Value {
		Insert(*current, &((*current).Left), value)
	} else {
		Insert(*current, &((*current).Right), value)
	}
}

func (t *BinarySearchTree) Search(value int) (int, error) {
	if t == nil {
		return 0, errors.New("No value was found")
	}

	if t.Value == value {
		return t.Value, nil
	}

	if value < t.Value {
		return t.Left.Search(value)
	} else {
		return t.Right.Search(value)
	}
}

func (t *BinarySearchTree) Min() *BinarySearchTree {
	if t.Left == nil {
		return t
	}

	return t.Left.Min()
}

func (t *BinarySearchTree) Max() *BinarySearchTree {
	if t.Right == nil {
		return t
	}

	return t.Right.Max()
}

func (t *BinarySearchTree) Traverse() []int {
	var values []int
	Traverse(t, &values)

	return values
}

func Traverse(bst *BinarySearchTree, list *[]int) {
	if bst == nil {
		return
	}

	Traverse(bst.Left, list)
	*list = append(*list, bst.Value)
	Traverse(bst.Right, list)
}

func (t *BinarySearchTree) Height() int {
	if t == nil {
		return 0
	}

	return Max(t.Left.Height(), t.Right.Height()) + 1
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
