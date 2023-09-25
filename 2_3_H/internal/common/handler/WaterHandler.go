package handler

import (
	"cosmos.collision/internal/common/sprite"
	"cosmos.collision/internal/utils"
	"reflect"
)

type WaterHandler struct {
	CollisionHandler
	nextHandler CollisionHandler
	typeOf      string
}

func NewWaterHandler(nextHandler CollisionHandler) CollisionHandler {
	return &WaterHandler{
		nextHandler: nextHandler,
		typeOf:      "*sprite.Water",
	}
}

func (w *WaterHandler) Handle(spritesMap map[int]interface{}, from int, to int) {
	adapter := NewCollisionAdapter(w, w.nextHandler, w.typeOf)
	adapter.Handling(spritesMap, from, to)
}

func (w *WaterHandler) Collision(_ sprite.ISprite, _ map[int]interface{}, toSprite sprite.ISprite) bool {
	if reflect.TypeOf(toSprite).String() == sprite.FireSprite {
		utils.MsgPrint(utils.DataWaterFire)
		return true
	}
	toSprite.(*sprite.Hero).AddHp()
	return false
}
