package common

import "fmt"

type ValueType interface {
	int | string
}

type IString[T ValueType] interface {
	Search(messages []string) T
	updateSearchResult(a string, b int) T
	searchEnd(index int, messages []string) bool
}

type StringHandler[T ValueType] struct {
	Param   T
	Handler IString[T]
}

func NewStringHandler[T ValueType](h IString[T]) *StringHandler[T] {
	return &StringHandler[T]{
		Handler: h,
	}
}

func (s *StringHandler[T]) Search(messages []string) T {
	for i, message := range messages {
		s.Param = s.Handler.updateSearchResult(message, i)
		fmt.Println("print:", message)
		if s.Handler.searchEnd(i, messages) {
			break
		}
	}
	return s.Param
}

func (s *StringHandler[T]) searchEnd(int, []string) bool {
	return false
}
