package common

import "fmt"

var level int

type Fan struct {
}

func init() {
	level = 0
}

type IFan interface {
	NextLevel()
	PreviousLevel()
}

func (f *Fan) NextLevel() {
	fmt.Println(fmt.Sprintf("[Fan] Next Level %d -> %d", level, level+1))
	level += 1
}

func (f *Fan) PreviousLevel() {
	fmt.Println(fmt.Sprintf("[Fan] Previous Level %d -> %d", level, level-1))
	level -= 1
}

type FanNextLevelCommand struct {
	fan *Fan
}

func NewFanNextLevelCommand(fan *Fan) ICommand {
	return &FanNextLevelCommand{fan: fan}
}

func (f *FanNextLevelCommand) Execute() {
	f.fan.NextLevel()
}

func (f *FanNextLevelCommand) Undo() {
	f.fan.PreviousLevel()
}

type FanPreviousLevelCommand struct {
	fan *Fan
}

func NewFanPreviousLevelCommand(fan *Fan) ICommand {
	return &FanPreviousLevelCommand{fan: fan}
}

func (f *FanPreviousLevelCommand) Execute() {
	f.fan.PreviousLevel()
}

func (f *FanPreviousLevelCommand) Undo() {
	f.fan.NextLevel()
}
