package handler

import (
	"cosmos.collision/internal/common/sprite"
	"reflect"
)

type HeroHandler struct {
	CollisionHandler
	nextHandler CollisionHandler
	typeOf      string
}

func NewHeroHandler(nextHandler CollisionHandler) CollisionHandler {
	return &HeroHandler{
		nextHandler: nextHandler,
		typeOf:      "*sprite.Hero",
	}
}

func (h *HeroHandler) Handle(spritesMap map[int]interface{}, from int, to int) {
	adapter := NewCollisionAdapter(h, h.nextHandler, h.typeOf)
	adapter.Handling(spritesMap, from, to)
}

func (h *HeroHandler) Collision(hero sprite.ISprite, spritesMap map[int]interface{}, toSprite sprite.ISprite) (isDead bool) {
	isDead = false
	if reflect.TypeOf(toSprite).String() == "*sprite.Water" {
		hero.(*sprite.Hero).AddHp()
	} else if reflect.TypeOf(toSprite).String() == "*sprite.Fire" {
		isDead = hero.(*sprite.Hero).MinusHp()
	}
	if !isDead {
		spritesMap[toSprite.GetPosition()] = hero
	}
	return
}
