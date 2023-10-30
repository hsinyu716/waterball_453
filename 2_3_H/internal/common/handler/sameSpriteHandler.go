package handler

import (
	"cosmos.collision/internal/common/sprite"
	"cosmos.collision/internal/utils"
	"reflect"
)

type SameSpriteHandler struct {
	ICollisionHandler
	nextHandler ICollisionHandler
}

func NewSameSpriteHandler(nextHandler ICollisionHandler) ICollisionHandler {
	return &SameSpriteHandler{
		nextHandler: nextHandler,
	}
}

func (s *SameSpriteHandler) Handle(spritePositions []sprite.ISprite, from, to int) {
	handler := NewCollisionHandler(s, s.nextHandler)
	handler.Handle(spritePositions, from, to)
}

func (s *SameSpriteHandler) match(fromSprite, toSprite sprite.ISprite) bool {
	return reflect.TypeOf(fromSprite) == reflect.TypeOf(toSprite)
}

func (s *SameSpriteHandler) Collision(_, _ sprite.ISprite, _ *[]sprite.ISprite) {
	utils.MsgPrint(utils.DataSameType)
}
