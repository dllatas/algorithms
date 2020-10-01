package balanced_parentheses

func naive(input string) bool {
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
