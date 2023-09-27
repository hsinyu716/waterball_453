package handler

import (
	"cosmos.collision/internal/common/sprite"
	"cosmos.collision/internal/utils"
	"reflect"
)

type BothDisappearHandler struct {
	ICollisionHandler
	nextHandler ICollisionHandler
}

func NewBothDisappearHandler(nextHandler ICollisionHandler) ICollisionHandler {
	return &BothDisappearHandler{
		nextHandler: nextHandler,
	}
}

func (f *BothDisappearHandler) Handle(spritePositions []sprite.ISprite, from, to int) {
	handler := NewCollisionHandler(f, f.nextHandler)
	handler.Handle(spritePositions, from, to)
}

func (f *BothDisappearHandler) match(_ *CollisionHandler, fromSprite, toSprite sprite.ISprite) bool {
	isMatch := false
	for _, fromType := range sprite.ConflictType {
		for _, toType := range sprite.ConflictType {
			if fromType != toType && fromType == reflect.TypeOf(fromSprite).String() && toType == reflect.TypeOf(toSprite).String() {
				isMatch = true
				break
			}
			if isMatch {
				break
			}
		}
	}
	return isMatch
}

func (f *BothDisappearHandler) Collision(formSprite, toSprite sprite.ISprite, spritePositions *[]sprite.ISprite) {
	formSprite.Remove(spritePositions)
	toSprite.Remove(spritePositions)
	utils.MsgPrint(utils.DataConflict)
}
