package handler

import (
	"cosmos.collision/internal/common/sprite"
	"cosmos.collision/internal/utils"
	"reflect"
)

type WaterHandler struct {
	CollisionHandler
	nextHandler CollisionHandler
	typeOf      sprite.TypeSprite
}

func NewWaterHandler(nextHandler CollisionHandler) CollisionHandler {
	return &WaterHandler{
		nextHandler: nextHandler,
		typeOf:      sprite.WaterSprite,
	}
}

func (w *WaterHandler) Handle(spritePositions []sprite.ISprite, from, to int) {
	adapter := NewCollisionAdapter(w, w.nextHandler, w.typeOf)
	adapter.Handling(spritePositions, from, to)
}

func (w *WaterHandler) Collision(_, toSprite sprite.ISprite, _ []sprite.ISprite) bool {
	if reflect.TypeOf(toSprite).String() == string(sprite.FireSprite) {
		utils.MsgPrint(utils.DataWaterFire)
		return true
	} else if reflect.TypeOf(toSprite).String() == string(sprite.IceSprite) {
		utils.MsgPrint(utils.DataIceFire)
		return true
	}
	toSprite.(*sprite.Hero).AddHp()
	return false
}
