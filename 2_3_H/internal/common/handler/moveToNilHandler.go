package handler

import (
	"cosmos.collision/internal/common/sprite"
	"cosmos.collision/internal/utils"
	"reflect"
)

type MoveToNilHandler struct {
	ICollisionHandler
	nextHandler ICollisionHandler
}

func NewMoveToNilHandler(nextHandler ICollisionHandler) ICollisionHandler {
	return &MoveToNilHandler{
		nextHandler: nextHandler,
	}
}

func (m *MoveToNilHandler) Handle(spritePositions []sprite.ISprite, from, to int) {
	handler := NewCollisionHandler(m, m.nextHandler)
	handler.Handle(spritePositions, from, to)
}

func (m *MoveToNilHandler) match(fromSprite, toSprite sprite.ISprite) bool {
	return fromSprite != nil && reflect.TypeOf(toSprite).String() == string(sprite.NilSprite)
}

func (m *MoveToNilHandler) Collision(fromSprite, toSprite sprite.ISprite, spritePositions *[]sprite.ISprite) {
	fromSprite.Move(toSprite, spritePositions)
	utils.MsgPrint(utils.DataMoveToNil)
}
