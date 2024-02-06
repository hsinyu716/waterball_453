package common

import (
	"fmt"
)

type Command struct{}

type ICommand interface {
	Execute()
	Undo()
}

type CommandCtrl struct{}

var (
	commands = make(map[int]ICommand, 6)
	s1       []ICommand
	s2       []ICommand
)

type ICommandCtrl interface {
	SetCommand(button int, command ICommand)
	Undo()
	Redo()
}

func (c *CommandCtrl) SetCommand(button int, command ICommand) {
	commands[button] = command
}

func (c *CommandCtrl) isEmpty(s []ICommand) bool {
	return len(s) == 0
}

func (c *CommandCtrl) pop(s []ICommand) (ICommand, []ICommand) {
	previousCommand := s[len(s)-1]
	s = s[:len(s)-1]
	return previousCommand, s
}

func (c *CommandCtrl) Undo() {
	if c.isEmpty(s1) {
		fmt.Println("s1 已無指令")
		return
	}
	previousCommand, s0 := c.pop(s1)
	s1 = s0
	previousCommand.Undo()
	s2 = append(s2, previousCommand)
}

func (c *CommandCtrl) Redo() {
	if c.isEmpty(s2) {
		fmt.Println("s2 已無指令")
		return
	}
	nextCommand, s0 := c.pop(s2)
	s2 = s0
	nextCommand.Execute()
	s1 = append(s1, nextCommand)
}

func (c *CommandCtrl) Press(button int) {
	if command, ok := commands[button]; ok {
		command.Execute()
		s1 = append(s1, command)
		s2 = nil
	} else {
		fmt.Println(fmt.Sprintf("%d unsupported", button))
	}
}
