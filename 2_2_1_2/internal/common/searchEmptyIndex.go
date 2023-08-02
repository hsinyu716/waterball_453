package common

type SearchEmpty[T ValueType] struct {
	StringHandler[T]
}

func (s *SearchEmpty[T]) updateSearchResult(a string, b int) T {
	if len(a) == 0 {
		return T(b)
	}
	b = -1
	return T(b)
}

func (s *SearchEmpty[T]) searchEnd(index int, messages []string) bool {
	return len(messages[index]) == 0
}
