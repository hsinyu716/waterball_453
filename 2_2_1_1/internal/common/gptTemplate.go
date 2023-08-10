package common

// Sorter 是排序的抽象父類別
type Sorter interface {
	Less(a, b int) bool
}

// SortBaseGPT 是 Sorter 接口的基本實現
type SortBaseGPT struct {
	Sorter Sorter
}

// Sort 是樣版方法，用來執行排序操作
func (s *SortBaseGPT) Sort(u []int) []int {
	n := len(u)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			// 使用注入的 Sorter 來比較
			if s.Sorter.Less(u[j], u[j+1]) {
				u[j], u[j+1] = u[j+1], u[j]
			}
		}
	}
	return u
}

// AscendingSort 是具體的子類別，實現升序排序
type AscendingSort struct {
	SortBaseGPT
}

// Less 是 AscendingSort 的 Less 實現，用來進行升序排序
func (a *AscendingSort) Less(x, y int) bool {
	return x > y
}

// DescendingSort 是具體的子類別，實現降序排序
type DescendingSort struct {
	SortBaseGPT
}

// Less 是 DescendingSort 的 Less 實現，用來進行降序排序
func (d *DescendingSort) Less(x, y int) bool {
	return x < y
}
