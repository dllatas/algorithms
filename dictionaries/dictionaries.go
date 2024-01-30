package dictionaries

import (
	"fmt"
)

type set struct {
	key   string
	value int
}

type dict struct {
	size   int
	values []*set
}

func newSet(key string, value int) *set {
	return &set{
		key:   key,
		value: value,
	}
}

func newDict() *dict {
	return &dict{}
}

// insert
func (d *dict) insert(key string, value int) {
	d.size += 1

	var set = newSet(key, value)
	d.values = append(d.values, set)
}

// get
func (d *dict) get(key string) (int, error) {
	for _, v := range d.values {
		if v.key == key {
			return v.value, nil
		}
	}
	return 0, fmt.Errorf("key %s not found.", key)
}

// remove
func (d *dict) remove(key string) {}
