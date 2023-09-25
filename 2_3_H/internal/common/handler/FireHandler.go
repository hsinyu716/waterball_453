package handler

import (
	"cosmos.collision/internal/common/sprite"
	"cosmos.collision/internal/utils"
	"reflect"
)

type FireHandler struct {
	CollisionHandler
	nextHandler CollisionHandler
	typeOf      string
}

func NewFireHandler(nextHandler CollisionHandler) CollisionHandler {
	return &FireHandler{
		nextHandler: nextHandler,
		typeOf:      "*sprite.Fire",
	}
}

func (f *FireHandler) Handle(spritesMap map[int]interface{}, from int, to int) {
	adapter := NewCollisionAdapter(f, f.nextHandler, f.typeOf)
	adapter.Handling(spritesMap, from, to)
}

func (f *FireHandler) Collision(_ sprite.ISprite, _ map[int]interface{}, toSprite sprite.ISprite) (isDead bool) {
	if reflect.TypeOf(toSprite).String() == sprite.WaterSprite {
		utils.MsgPrint(utils.DataWaterFire)
		return true
	}
	isDead = toSprite.(*sprite.Hero).MinusHp()
	return
}
