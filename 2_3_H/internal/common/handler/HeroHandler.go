package handler

import (
	"cosmos.collision/internal/common/sprite"
	"reflect"
)

type HeroHandler struct {
	CollisionHandler
	nextHandler CollisionHandler
	typeOf      sprite.TypeSprite
}

func NewHeroHandler(nextHandler CollisionHandler) CollisionHandler {
	return &HeroHandler{
		nextHandler: nextHandler,
		typeOf:      sprite.HeroSprite,
	}
}

func (h *HeroHandler) Handle(spritesMap map[int]interface{}, from, to int) {
	adapter := NewCollisionAdapter(h, h.nextHandler, h.typeOf)
	adapter.Handling(spritesMap, from, to)
}

func (h *HeroHandler) Collision(hero, toSprite sprite.ISprite, spritesMap map[int]interface{}) (isDead bool) {
	isDead = false
	switch reflect.TypeOf(toSprite).String() {
	case string(sprite.WaterSprite):
		fallthrough
	case string(sprite.IceSprite):
		hero.(*sprite.Hero).AddHp()
		break
	case string(sprite.FireSprite):
		isDead = hero.(*sprite.Hero).MinusHp()
		break
	}
	if !isDead {
		spritesMap[toSprite.GetPosition()] = hero
	}
	return
}
