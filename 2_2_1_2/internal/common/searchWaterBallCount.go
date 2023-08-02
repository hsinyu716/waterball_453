package common

import "strings"

type SearchCount[T ValueType] struct {
	StringHandler[T]
	count int
}

func (s *SearchCount[T]) updateSearchResult(a string, b int) T {
	if strings.EqualFold(a, "WaterBall") {
		s.count++
	}
	return T(s.count)
}
