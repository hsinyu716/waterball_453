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
	ctrl := NewCollisionHandler(h, h.nextHandler)
	ctrl.Handle(spritePositions, from, to)
}

func (h *HeroStrengthenHandler) match(fromSprite, toSprite sprite.ISprite) bool {
	match := false
	for _, hpType := range sprite.StrengthenType {
		match = reflect.TypeOf(fromSprite).String() == string(sprite.HeroSprite) && reflect.TypeOf(toSprite).String() == hpType ||
			reflect.TypeOf(fromSprite).String() == hpType && reflect.TypeOf(toSprite).String() == string(sprite.HeroSprite)
		if match {
			break
		}
	}
	return match
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
