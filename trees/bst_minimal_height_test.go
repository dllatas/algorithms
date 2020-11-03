package trees

import (
	"log"
	"math"
	"sort"
	"testing"
	"time"
)

type TestMinHeightBST struct {
	data []int
	size int
	want int
	bst  *BinarySearchTree
}

func NewTestMinHeightBST(inputList [][]int) []TestMinHeightBST {
	var tests []TestMinHeightBST

	var unique = func(input []int) []int {
		keys := make(map[int]bool)
		list := []int{}

		for _, value := range input {
			if _, found := keys[value]; !found {
				keys[value] = true
				list = append(list, value)
			}
		}
		return list
	}

	for _, input := range inputList {
		var noDups = unique(input)

		// be aware that sequentially inserting
		// a sorted array makes the tree unbalanced
		// unbalanced tree -> no min height
		sort.Ints(noDups)

		var size = len(noDups)

		// 2 to the power of n = len(fixture), n = levels
		var height = int(math.Logb(float64(size))) + 1
		if size == 0 {
			height = 0
		}

		var test = TestMinHeightBST{
			data: noDups,
			size: size,
			want: height,
		}

		tests = append(tests, test)
	}

	return tests
}

var TestMinHeightBSTFixtures = [][]int{
	{},
	{5},
	{34, 568},
	{8, 8, 8},
	{56, 78, 23, 45},
	{34, 2, 42, 56, 78, 12, 27, 5, 8},
	{1, 8, 3, 6, 5, 4, 7, 2, 9, 0, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10},
	{72, 14, 78, 9, 90, 8, 91, 93, 56, 66, 3, 92, 47, 16, 22, 7, 40, 75, 11, 46, 100, 89, 24, 85, 74, 58, 36, 86, 60, 99, 70, 84, 83, 59, 52, 38, 43, 64, 25, 31, 1, 48, 79, 49, 17, 73, 15, 76, 95, 53, 18, 29, 2, 45, 77, 13, 20, 34, 97, 32, 54, 65, 41, 81, 39, 5, 96, 68, 82, 55, 19, 50, 87, 44, 80, 98, 30, 6, 23, 4, 51, 42, 21, 26, 71, 62, 27, 12, 10, 88, 61, 69, 63, 67, 28, 37, 35, 57, 33, 94},
}

var testsMinHeightBST = NewTestMinHeightBST(TestMinHeightBSTFixtures)

func TestSimpleMinHeightBST(t *testing.T) {
	start := time.Now()
	for _, test := range testsMinHeightBST {
		var bst *BinarySearchTree

		simpleMinHeightBST(&bst, test.data, test.size)
		var got = bst.Height()

		if test.want != got {
			t.Errorf("got: %v, want: %v", got, test.want)
		}
	}
	duration := time.Since(start)
	log.Printf("Simple Min Height BST: %s", duration)
}

func TestDirectMinHeightBST(t *testing.T) {
	start := time.Now()
	for _, test := range testsMinHeightBST {
		var bst *BinarySearchTree

		directMinHeightBST(&bst, &bst, test.data, test.size)
		var got = bst.Height()

		if test.want != got {
			t.Errorf("got: %v, want: %v", got, test.want)
		}
	}
	duration := time.Since(start)
	log.Printf("Direct Min Height BST: %s", duration)
}
