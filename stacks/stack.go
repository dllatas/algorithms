package stack

// Element ...
type Element struct {
	value    string
	hasValue bool
	position int
}

// NewElement ...
func NewElement(value string, position int) *Element {
	var el Element
	el = Element{
		value:    value,
		position: position,
		hasValue: true,
	}
	if value == "" {
		el.position = 0
		el.hasValue = false
	}
	return &el
}

// Stack ...
type Stack struct {
	values []*Element
	size   int
}

// NewStack ...
func NewStack() *Stack {
	var v []*Element
	var stack = Stack{
		values: v,
		size:   0,
	}
	return &stack
}

func (s *Stack) push(el *Element) {
	s.values = append(s.values, el)
	s.size = s.size + 1
}

func (s *Stack) add(value string, position int) {
	var el = NewElement(value, position)

	s.values = append(s.values, el)
	s.size = s.size + 1
}

func (s *Stack) pop() *Element {
	if s.size == 0 {
		return NewElement("", s.size)
	}

	var lastIndex = s.size - 1
	var last = s.values[lastIndex]

	s.values = s.values[0:lastIndex]
	s.size = s.size - 1

	return last
}

func (s *Stack) popo() {
	if s.size == 0 {
		return
	}

	var lastIndex = s.size - 1
	s.values = s.values[0:lastIndex]
	s.size = s.size - 1
}

func (s *Stack) last() *Element {
	if s.size == 0 {
		return NewElement("", s.size)
	}

	var lastIndex = s.size - 1
	var last = s.values[lastIndex]

	return last
}
