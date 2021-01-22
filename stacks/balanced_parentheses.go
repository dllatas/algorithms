package stack

func balance(input string) bool {
	var right = '('
	var left = ')'
	var counterRight = 0
	var counterLeft = 0
	for _, value := range input {
		if value == right {
			counterRight = counterRight + 1
		}
		if value == left {
			counterLeft = counterLeft + 1
		}
		if counterLeft > counterRight {
			return false
		}

	}
	if counterLeft != counterRight {
		return false
	}
	return true
}

func balanceUsingStack(input string) (bool, int) {
	var right = "("
	var left = ")"

	var stack = NewStack()

	for index, value := range input {
		var el = NewElement(string(value), index)
		if el.value == right {
			stack.push(el)
		}

		if el.value == left {
			var current = stack.pop()
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
