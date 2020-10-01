package balanced_parentheses

type Element struct {
	value    rune
	position int
	hasValue bool
}

type Stack struct {
	values []Element
	size   int
}

func NewStack() *Stack {
	var v []Element
	var stack = Stack{
		values: v,
		size:   0,
	}
	return &stack
}

func (s *Stack) Push(el Element) {
	s.values = append(s.values, el)
	s.size = s.size + 1
}

func (s *Stack) Pop() Element {
	if s.size == 0 {
		return Element{
			hasValue: false,
		}
	}

	var last = s.values[s.size-1]
	s.values = s.values[0 : s.size-1]
	s.size = s.size - 1
	return last
}

func checkWithStack(input string) (bool, int) {
	var right = '('
	var left = ')'

	var stack = NewStack()

	for index, value := range input {
		var el = Element{
			value:    value,
			position: index,
			hasValue: true,
		}

		if el.value == right {
			stack.Push(el)
		}

		if el.value == left {
			var current = stack.Pop()
			if !current.hasValue {
				return false, index
			}
		}
	}

	if stack.size > 0 {
		return false, stack.values[0].position
	}

	return true, -1
}
