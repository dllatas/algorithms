package stack

type intElement struct {
	value    int
	position int
}

type stackInt struct {
	values []*intElement
	size   int
}

func newIntElement(value, position int) *intElement {
	var el intElement
	el = intElement{
		value:    value,
		position: position,
	}
	return &el
}

func newStackInt() *stackInt {
	var v []*intElement
	var stack = stackInt{
		values: v,
		size:   0,
	}
	return &stack
}

func (s *stackInt) push(value int) {
	var el = newIntElement(value, s.size)
	s.values = append(s.values, el)
	s.size = s.size + 1
}

func (s *stackInt) pop() *intElement {
	if s.size == 0 {
		return newIntElement(0, s.size)
	}

	var lastIndex = s.size - 1
	var last = s.values[lastIndex]

	s.values = s.values[0:lastIndex]
	s.size = s.size - 1

	return last
}

func (s *stackInt) peek() (*intElement, bool) {
	if s.size == 0 {
		// TODO: return error instead?
		return newIntElement(0, s.size), true
	}

	var lastIndex = s.size - 1
	var peek = s.values[lastIndex]

	return peek, false
}
