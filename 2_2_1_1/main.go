package main

import (
	"cosmos.template/internal/common"
	"fmt"
)

func main() {
	u := []int{7, 2, 5, 1, 8, 3}
	// 使用升序排序
	baseGPT := &common.SortBaseGPT{Sorter: &common.AscendingSort{}}
	fmt.Println("升序排序結果：", baseGPT.Sort(u))

	// 使用降序排序
	baseGPT = &common.SortBaseGPT{Sorter: &common.DescendingSort{}}
	fmt.Println("降序排序結果：", baseGPT.Sort(u))

	u = []int{7, 2, 5, 1, 8, 3}
	base := common.NewSortBase(&common.SortASC{})
	fmt.Println("升序排序結果：", base.Sort(u))

	base = common.NewSortBase(&common.SortDESC{})
	fmt.Println("降序排序結果：", base.Sort(u))
}
