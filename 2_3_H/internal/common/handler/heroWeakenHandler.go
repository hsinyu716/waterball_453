package handler

import (
	"cosmos.collision/internal/common/sprite"
	"reflect"
)

type HeroWeakenHandler struct {
	ICollisionHandler
	nextHandler ICollisionHandler
}

func NewHeroWeakenHandler(nextHandler ICollisionHandler) ICollisionHandler {
	return &HeroWeakenHandler{
		nextHandler: nextHandler,
	}
}

func (h *HeroWeakenHandler) Handle(spritePositions []sprite.ISprite, from, to int) {
	adapter := NewCollisionHandler(h, h.nextHandler)
	adapter.Handle(spritePositions, from, to)
}

func (h *HeroWeakenHandler) match(fromSprite, toSprite sprite.ISprite) bool {
	return sprite.StrengthenOrWeaken(fromSprite, toSprite, sprite.WeakenType)
}

func (h *HeroWeakenHandler) Collision(fromSprite, toSprite sprite.ISprite, spritePositions *[]sprite.ISprite) {
	isDead := false
	fromSprite.Remove(spritePositions)
	if reflect.TypeOf(fromSprite).String() == string(sprite.HeroSprite) {
		isDead = fromSprite.(*sprite.Hero).MinusHp()
		if !isDead {
			fromSprite.(*sprite.Hero).Move(toSprite, spritePositions)
		}
	} else {
		isDead = toSprite.(*sprite.Hero).MinusHp()
	}
	if isDead {
		toSprite.Remove(spritePositions)
	}
}
