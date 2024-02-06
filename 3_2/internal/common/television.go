package common

import "fmt"

type Television struct {
}

type ITelevision interface {
	TurnOn()
	TurnOff()
}

func (t *Television) TurnOn() {
	fmt.Println("[TV] TurnOn.")
}

func (t *Television) TurnOff() {
	fmt.Println("[TV] TurnOff.")
}

type TelevisionTurnOnCommand struct {
	tv *Television
}

func NewTVTurnOn(tv *Television) ICommand {
	return &TelevisionTurnOnCommand{tv: tv}
}

func (t *TelevisionTurnOnCommand) Execute() {
	t.tv.TurnOn()
}
func (t *TelevisionTurnOnCommand) Undo() {
	t.tv.TurnOff()
}

type TelevisionTurnOffCommand struct {
	tv *Television
}

func NewTVTurnOff(tv *Television) ICommand {
	return &TelevisionTurnOffCommand{tv: tv}
}

func (t *TelevisionTurnOffCommand) Execute() {
	t.tv.TurnOff()
}

func (t *TelevisionTurnOffCommand) Undo() {
	t.tv.TurnOn()
}
