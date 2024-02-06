package common

import "fmt"

type AirConditioner struct {
}

type IAirConditioner interface {
	ITelevision
}

func (a *AirConditioner) TurnOn() {
	fmt.Println("[AirConditioner] TurnOn.")
}

func (a *AirConditioner) TurnOff() {
	fmt.Println("[AirConditioner] TurnOff.")
}

type AirConditionerTurnOnCommand struct {
	ac *AirConditioner
}

func NewACTurnOn(ac *AirConditioner) ICommand {
	return &AirConditionerTurnOnCommand{ac: ac}
}

func (a *AirConditionerTurnOnCommand) Execute() {
	a.ac.TurnOn()
}

func (a *AirConditionerTurnOnCommand) Undo() {
	a.ac.TurnOff()
}

type AirConditionerTurnOffCommand struct {
	ac *AirConditioner
}

func NewACTurnOff(ac *AirConditioner) ICommand {
	return &AirConditionerTurnOffCommand{ac: ac}
}

func (a *AirConditionerTurnOffCommand) Execute() {
	a.ac.TurnOff()
}

func (a *AirConditionerTurnOffCommand) Undo() {
	a.ac.TurnOn()
}
