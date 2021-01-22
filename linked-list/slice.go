package linkedlist

func (l *List) sliceToLast(from int) (int, *List) {
	if from == 0 {
		return 0, &List{}
	}

	if l.Next == nil {
		return 1, l
	}

	index, value := l.Next.sliceToLast(from)
	if index < from {
		value = l
		index = index + 1
		return index, value
	}

	return index, value
}

func (l *List) length() int {
	if l.Next == nil {
		return 1
	}
	var length = l.Next.length()
	return length + 1
}
