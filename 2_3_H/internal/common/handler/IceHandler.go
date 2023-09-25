package handler

import (
	"cosmos.collision/internal/common/sprite"
	"cosmos.collision/internal/utils"
	"reflect"
)

type IceHandler struct {
	CollisionHandler
	nextHandler CollisionHandler
	typeOf      sprite.TypeSprite
}

func NewIceHandler(nextHandler CollisionHandler) CollisionHandler {
	return &IceHandler{
		nextHandler: nextHandler,
		typeOf:      sprite.IceSprite,
	}
}

func (f *IceHandler) Handle(spritePositions []sprite.ISprite, from, to int) {
	adapter := NewCollisionAdapter(f, f.nextHandler, f.typeOf)
	adapter.Handling(spritePositions, from, to)
}

func (f *IceHandler) Collision(_, toSprite sprite.ISprite, _ []sprite.ISprite) (isDead bool) {
	if reflect.TypeOf(toSprite).String() == string(sprite.WaterSprite) {
		utils.MsgPrint(utils.DataIceFire)
		return true
	} else if reflect.TypeOf(toSprite).String() == string(sprite.FireSprite) {
		utils.MsgPrint(utils.DataIceFire)
		return true
	}
	toSprite.(*sprite.Hero).AddHp()
	return false
}
