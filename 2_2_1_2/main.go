package main

import (
	"cosmos.template2/internal/common"
	"fmt"
)

func main() {
	c1 := common.SearchLongest[string]{}
	u := []string{
		"cat",
		"a",
		"already",
		"ab",
		"construct",
	}
	handler := common.NewStringHandler[string](&c1)
	str := handler.Search(u)
	fmt.Println("最長字串", str)

	c2 := common.SearchEmpty[int]{}
	u2 := []string{
		"a",
		"d",
		"",
		"ccc",
	}
	handler2 := common.NewStringHandler[int](&c2)
	str2 := handler2.Search(u2)
	fmt.Println("空字串索引", str2)

	c3 := common.SearchCount[int]{}
	u3 := []string{
		"WaterBall",
		"",
		"WaterBall",
		"ccc",
	}
	handler3 := common.NewStringHandler[int](&c3)
	str3 := handler3.Search(u3)
	fmt.Println("次數", str3)
}
