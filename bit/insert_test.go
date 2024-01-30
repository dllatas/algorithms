package bit

import (
	"testing"
)

type InsertTest struct {
	m    uint64
	n    uint64
	i    int
	j    int
	want uint64
}

func TestInsert(t *testing.T) {
	var test = InsertTest{
		n:    10000000000,
		m:    10011,
		i:    2,
		j:    6,
		want: 10001001100,
	}

	var got = Insert(test.n, test.m, test.i, test.j)

	if got != test.want {
		t.Errorf("got: %v, want: %v", got, test.want)
	}
}
