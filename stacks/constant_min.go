package stack

import (
	"log"
)

type stackMin struct {
	min    *stackInt
	values *stackInt
}

func newStackMin() *stackMin {
	var min = newStackInt()
	var values = newStackInt()
	var stackMin = stackMin{
		min:    min,
		values: values,
	}
	return &stackMin
}

func (s *stackMin) push(value int) {
	s.values.push(value)

	peek, isEmpty := s.min.peek()
	if isEmpty {
		s.min.push(value)
		return
	}

	if value <= peek.value {
		s.min.push(value)
		return
	}
}

func (s *stackMin) pop() {
	peek, isEmpty := s.values.peek()
	if isEmpty {
		return
	}

	s.values.pop()

	minPeek, _ := s.min.peek()
	if peek.value == minPeek.value {
		s.min.pop()
	}
}

func (s *stackMin) peek() (*intElement, bool) {
	peek, isEmpty := s.values.peek()
	return peek, isEmpty
}

func (s *stackMin) getMin() (*intElement, bool) {
	log.Printf("%v", s.min)
	minPeek, isEmpty := s.min.peek()
	return minPeek, isEmpty
}
