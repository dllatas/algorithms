package trees

import (
	"log"
	"testing"
	//	"github.com/dllatas/algorithms/list"
)

/*
type ListPerDepthTest struct {
	values []int
	tree   *BinarySearchTree
	want   map[int]*list.List
}
*/

func TestGetLists(t *testing.T) {
	//var values = []int{1, 2, 3, 4, 5, 6, 7}
	var values = []int{4, 2, 6, 1, 3, 5, 7}
	/*
		var want = map[int]*list.List{
			1: list.NewListFromSlice([]int{4}),
			2: list.NewListFromSlice([]int{2, 6}),
			3: list.NewListFromSlice([]int{1, 3, 5, 7}),
		}
	*/
	var tree = NewBinarySearchTree(values[0])

	for i := 1; i < len(values); i++ {
		tree.Insert(values[i])
	}

	//tree.describe()

	var lists = getLists(tree)

	log.Printf("%v", lists)
}
