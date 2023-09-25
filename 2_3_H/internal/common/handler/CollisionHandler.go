package handler

import (
	"cosmos.collision/internal/common/sprite"
	"fmt"
	"reflect"
)

type CollisionHandler interface {
	Handle(spritesMap map[int]interface{}, from int, to int)
	Collision(fromSprite sprite.ISprite, spritesMap map[int]interface{}, toSprite sprite.ISprite) bool
}

type CollisionAdapter struct {
	handler     CollisionHandler
	nextHandler CollisionHandler
	typeOf      string
}

func NewCollisionAdapter(handler CollisionHandler, nextHandler CollisionHandler, typeOf string) *CollisionAdapter {
	return &CollisionAdapter{
		handler:     handler,
		nextHandler: nextHandler,
		typeOf:      typeOf,
	}
}

func (c *CollisionAdapter) Handling(spritesMap map[int]interface{}, from int, to int) {
	fromSprite := spritesMap[from].(sprite.ISprite)
	toSprite := spritesMap[to].(sprite.ISprite)
	if spritesMap[to] == nil {
		// TODO:討論  可以放這  或放world
		spritesMap[to] = spritesMap[from]
		spritesMap[from] = nil
		return
	}
	if reflect.TypeOf(fromSprite).String() == c.typeOf {
		c.Collision(fromSprite, spritesMap, toSprite)
		spritesMap[from] = nil
	} else {
		c.next(c.nextHandler, spritesMap, from, to)
	}
}

func (c *CollisionAdapter) Collision(fromSprite sprite.ISprite, spritesMap map[int]interface{}, toSprite sprite.ISprite) bool {
	fmt.Println(fmt.Sprintf("--%v >> %v", reflect.TypeOf(fromSprite), reflect.TypeOf(toSprite)))
	isDead := false
	isDead = c.handler.Collision(fromSprite, spritesMap, toSprite)
	if isDead {
		toSprite.Remove(&spritesMap)
	}
	return isDead
}

func (c *CollisionAdapter) next(handler CollisionHandler, spritesMap map[int]interface{}, from int, to int) {
	if handler != nil {
		handler.Handle(spritesMap, from, to)
	}
}
