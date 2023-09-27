package handler

import (
	"cosmos.collision/internal/common/sprite"
	"reflect"
)

type HeroStrengthenHandler struct {
	ICollisionHandler
	nextHandler ICollisionHandler
}

func NewHeroStrengthenHandler(nextHandler ICollisionHandler) ICollisionHandler {
	return &HeroStrengthenHandler{
		nextHandler: nextHandler,
	}
}

func (h *HeroStrengthenHandler) Handle(spritePositions []sprite.ISprite, from, to int) {
	adapter := NewCollisionHandler(h, h.nextHandler)
	adapter.Handle(spritePositions, from, to)
}

func (h *HeroStrengthenHandler) match(handler *CollisionHandler, fromSprite, toSprite sprite.ISprite) bool {
	return handler.strengthenOrWeaken(fromSprite, toSprite, sprite.StrengthenType)
}

func (h *HeroStrengthenHandler) Collision(fromSprite, toSprite sprite.ISprite, spritePositions *[]sprite.ISprite) {
	fromSprite.Remove(spritePositions)
	for _, hpType := range sprite.StrengthenType {
		if reflect.TypeOf(toSprite).String() == hpType {
			fromSprite.(*sprite.Hero).AddHp()
			fromSprite.(*sprite.Hero).Move(toSprite, spritePositions)
		} else if reflect.TypeOf(fromSprite).String() == hpType {
			toSprite.(*sprite.Hero).AddHp()
		}
	}
}
