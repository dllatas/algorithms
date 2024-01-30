package stack

func mutipleBalance(input string, specs map[string]string) bool {
	var stack = NewStack()

	for index, value := range input {
		if ok := isValid(string(value), specs); !ok {
			continue
		}

		_, found := specs[string(value)]
		if found {
			stack.add(string(value), index)
		}

		if !found {
			var last = stack.last()
			var closing = specs[last.value]

			if closing == string(value) {
				stack.popo()
			} else {
				return false
			}
		}
	}

	if stack.size > 0 {
		return false
	}

	return true
}

func isValid(s string, specs map[string]string) bool {
	_, found := specs[s]
	if found {
		return true
	}

	for _, v := range specs {
		if v == s {
			return true
		}
	}

	return false
}
