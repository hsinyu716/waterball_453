package main

import (
	"cosmos.collision/internal/common/world"
	"fmt"
)

func main() {
	var x int
	var y int

	w := world.World{}
	w.Init()

	for {
		fmt.Println("請輸入二個整數0~29，使用空白隔開：")
		_, err := fmt.Scanln(&x, &y)
		if err != nil {
			fmt.Println("請輸入二個值")
			x = -1
			y = -1
		}
		if check(x) && check(y) {
			w.Move(x, y)
		} else {
			fmt.Println("值區錯誤")
		}
	}
}

func check(v int) bool {
	return v >= 0 && v <= 29
}
