package main

import (
	"cosmos.appliances/internal/common"
	"fmt"
)

func main() {
	fan := new(common.Fan)
	tv_ := new(common.Television)
	ac_ := new(common.AirConditioner)
	ctrl := new(common.CommandCtrl)
	ctrl.SetCommand(0, common.NewFanNextLevelCommand(fan))
	ctrl.SetCommand(1, common.NewFanPreviousLevelCommand(fan))
	ctrl.SetCommand(2, common.NewTVTurnOn(tv_))
	ctrl.SetCommand(3, common.NewTVTurnOff(tv_))
	ctrl.SetCommand(4, common.NewACTurnOn(ac_))
	ctrl.SetCommand(5, common.NewACTurnOff(ac_))

	for {
		command := -3
		fmt.Println("請輸入指令[0-5], or undo(-1) redo(-2)")
		_, _ = fmt.Scan(&command)
		if command == -1 {
			ctrl.Undo()
		} else if command == -2 {
			ctrl.Redo()
		} else {
			ctrl.Press(command)
		}
	}
}
