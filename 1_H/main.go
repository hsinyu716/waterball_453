package main

import (
	"waterball/1_H/internal/common"
)

func main() {
	p1 := new(common.Human)
	p2 := new(common.AI)
	p3 := new(common.Human)
	p4 := new(common.Human)

	desk := common.NewDesk()
	showdown := common.NewShowdown(desk, &[]common.PlayerService{p1, p2, p3, p4})
	showdown.Start()
}
