package handler

import (
	"cosmos.collision/internal/common/sprite"
	"cosmos.collision/internal/utils"
	"reflect"
)

type FireHandler struct {
	CollisionHandler
	nextHandler CollisionHandler
	typeOf      sprite.TypeSprite
}

func NewFireHandler(nextHandler CollisionHandler) CollisionHandler {
	return &FireHandler{
		nextHandler: nextHandler,
		typeOf:      sprite.FireSprite,
	}
}

func (f *FireHandler) Handle(spritesMap map[int]interface{}, from, to int) {
	adapter := NewCollisionAdapter(f, f.nextHandler, f.typeOf)
	adapter.Handling(spritesMap, from, to)
}

func (f *FireHandler) Collision(_, toSprite sprite.ISprite, _ map[int]interface{}) (isDead bool) {
	if reflect.TypeOf(toSprite).String() == string(sprite.WaterSprite) {
		utils.MsgPrint(utils.DataWaterFire)
		return true
	} else if reflect.TypeOf(toSprite).String() == string(sprite.IceSprite) {
		utils.MsgPrint(utils.DataWaterFire)
		return true
	}
	isDead = toSprite.(*sprite.Hero).MinusHp()
	return
}
