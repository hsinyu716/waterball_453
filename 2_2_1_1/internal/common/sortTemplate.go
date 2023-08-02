package common

type ISort interface {
	Sort(u []int) []int
	Compare(a, b int) bool
}

type SortBase struct {
	iSort ISort
}

func NewSortBase(iS ISort) ISort {
	return &SortBase{
		iSort: iS,
	}
}

func (s *SortBase) Sort(u []int) []int {
	n := len(u)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if s.iSort.Compare(u[j], u[j+1]) {
				u[j], u[j+1] = u[j+1], u[j]
			}
		}
	}
	return u
}

func (s *SortBase) Compare(a, b int) bool {
	return true
}

type SortASC struct {
	SortBase
}

func (s *SortASC) Compare(a, b int) bool {
	return a > b
}

type SortDESC struct {
	SortBase
}

func (s *SortDESC) Compare(a, b int) bool {
	return a < b
}
