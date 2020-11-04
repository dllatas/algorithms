package sorting

type SortRunes []rune

func (s SortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortRunes) Len() int {
	return len(s)
}

func (s SortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
