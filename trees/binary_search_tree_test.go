package trees

import (
	//"log"
	"reflect"
	"sort"
	"testing"
)

type BinarySearchTreeTests struct {
	Value int
	Want  *BinarySearchTree
}

func TestNewBinarySearchTree(t *testing.T) {
	var tests = []BinarySearchTreeTests{
		{
			Value: 5,
			Want: &BinarySearchTree{
				Value: 5,
			},
		},
	}

	for _, test := range tests {
		var got = NewBinarySearchTree(test.Value)
		if !reflect.DeepEqual(got, test.Want) {
			t.Errorf("got: %v, want: %v", got, test.Want)
		}
	}
}

func TestInsert(t *testing.T) {
	var fixture = []int{1, 2, 3, 4, 5}
	var bst = NewBinarySearchTree(fixture[0])

	if bst.Value != fixture[0] {
		t.Errorf("got: %v, want: %v", bst.Value, fixture[0])
	}

	bst.Insert(fixture[1])
	if bst.Right.Value != fixture[1] {
		t.Errorf("got: %v, want: %v", bst.Right.Value, fixture[1])
	}

	bst.Insert(fixture[2])
	if bst.Right.Right.Value != fixture[2] {
		t.Errorf("got: %v, want: %v", bst.Right.Value, fixture[2])
	}

	bst.Insert(fixture[3])
	if bst.Right.Right.Right.Value != fixture[3] {
		t.Errorf("got: %v, want: %v", bst.Right.Right.Value, fixture[3])
	}
}

func TestSearch(t *testing.T) {
	var fixture = []int{8, 3, 1, 9}

	var bst = NewBinarySearchTree(5)
	for _, value := range fixture {
		bst.Insert(value)
		found, ok := bst.Search(value)
		if ok != nil {
			t.Errorf("%v", ok)
		}

		if found != value {
			t.Errorf("got: %v, want: %v", found, value)
		}
	}

	_, ok := bst.Search(12)
	if ok == nil {
		t.Errorf("%v", ok)
	}
}

func TestMin(t *testing.T) {
	var fixture = []int{8, 3, 1, 9, 6, 34, 45, 67}

	var bst = NewBinarySearchTree(5)
	for _, value := range fixture {
		bst.Insert(value)
	}

	var min = bst.Min()
	if min.Value != 1 {
		t.Errorf("got: %v, want: %v", min.Value, 1)
	}

	if min.Parent.Value != 3 {
		t.Errorf("got: %v, want: %v", min.Parent.Value, 3)
	}
}

func TestMax(t *testing.T) {
	var fixture = []int{8, 3, 1, 9, 6, 34, 45, 67}

	var bst = NewBinarySearchTree(5)
	for _, value := range fixture {
		bst.Insert(value)
	}

	var max = bst.Max()
	if max.Value != 67 {
		t.Errorf("got: %v, want: %v", max.Value, 67)
	}

	if max.Parent.Value != 45 {
		t.Errorf("got: %v, want: %v", max.Parent.Value, 45)
	}
}

func TestTraverse(t *testing.T) {
	var root = 5
	var fixture = []int{34, 45, 8, 3, 1, 9, 6, 67}

	var bst = NewBinarySearchTree(root)
	for _, value := range fixture {
		bst.Insert(value)
	}

	var got = bst.Traverse()
	fixture = append(fixture, root)
	sort.Ints(fixture)

	if !reflect.DeepEqual(got, fixture) {
		t.Errorf("got: %v, want: %v", got, fixture)
	}
}
