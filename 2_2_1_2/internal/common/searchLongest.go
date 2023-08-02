package common

type SearchLongest[T ValueType] struct {
	StringHandler[T]
	maxLengthMessage interface{}
}

func (s *SearchLongest[T]) updateSearchResult(a string, b int) T {
	s.maxLengthMessage = ""
	if len(a) > len(s.maxLengthMessage.(string)) {
		s.maxLengthMessage = a
	}
	return s.maxLengthMessage.(T)
}
