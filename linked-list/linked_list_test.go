package linked_list

import (
	//"log"
	"reflect"
	"testing"
)

/*func TestInsert(t *testing.T) {
	for _, tt := range tests {
		got := naive(tt.input)

		if !(got == tt.want) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}*/

type Tests []struct {
	input   int
	want    []int
	current int
}

func TestInsert(t *testing.T) {
	var got []int
	var tests = Tests{
		{
			input:   5,
			want:    []int{5},
			current: 5,
		},
		{
			input:   6,
			want:    []int{5, 6},
			current: 6,
		},
		{
			input:   7,
			want:    []int{5, 6, 7},
			current: 7,
		},
	}

	var list = NewList(tests[0].input)
	got = append(got, list.Value)
	if !reflect.DeepEqual(got, tests[0].want) {
		t.Errorf("got %v, want %v", got, tests[0].want)
	}

	for i := 1; i < len(tests); i++ {
		list = Insert(list, tests[i].input)
		if list.Value != tests[i].current {
			t.Errorf("got %v, want %v", got, tests[i].want)
		}

		got = append(got, list.Value)

		if !reflect.DeepEqual(got, tests[i].want) {
			t.Errorf("got %v, want %v", got, tests[i].want)
		}
	}
}

func TestSearch(t *testing.T) {
	var tests = []int{6, 34, 10, 55, 8}

	var list = NewList(tests[0])
	for i := 1; i < len(tests); i++ {
		list = Insert(list, tests[i])
	}

	for i := len(tests) - 1; i >= 0; i-- {
		var got = Search(list, tests[i])
		if got.Value != tests[i] {
			t.Errorf("got %v, want %v", got.Value, tests[i])
		}
	}

	// Special case: Element not in list
	var got = Search(list, 99)
	if got != nil {
		t.Errorf("got %v, want %v", got, nil)
	}
}

func TestGetPrevious(t *testing.T) {
	var tests = []int{6, 34, 10, 55, 8}

	var list = NewList(tests[0])
	for i := 1; i < len(tests); i++ {
		list = Insert(list, tests[i])
	}

	for i := 0; i < len(tests)-1; i++ {
		var prevList = GetPrevious(list, tests[i])
		if prevList.Value != tests[i+1] {
			t.Errorf("got %v, want %v", prevList.Value, tests[i-1])
		}
	}

	// Special case: Prev for first element
	var prevList = GetPrevious(list, tests[len(tests)-1])
	if prevList != nil {
		t.Errorf("got %v, want %v", prevList, nil)
	}
}

func TestDelete(t *testing.T) {
	var tests = []int{6, 34, 10, 55, 8}

	var list = NewList(tests[0])
	for i := 1; i < len(tests); i++ {
		list = Insert(list, tests[i])
	}

	var itemsToRemove = []int{1, 3, 4}
	for _, value := range itemsToRemove {
		Delete(list, tests[value])
		var found = Search(list, tests[value])
		if found != nil {
			t.Errorf("got %v, want %v", found, nil)
		}
	}
}