package handler

import (
	"cosmos.collision/internal/common/sprite"
	"fmt"
	"reflect"
)

type CollisionHandler interface {
	Handle(spritePositions []sprite.ISprite, from, to int)
	Collision(fromSprite, toSprite sprite.ISprite, spritePositions []sprite.ISprite) bool
}

type CollisionAdapter struct {
	handler     CollisionHandler
	nextHandler CollisionHandler
	typeOf      sprite.TypeSprite
}

func NewCollisionAdapter(handler, nextHandler CollisionHandler, typeOf sprite.TypeSprite) *CollisionAdapter {
	return &CollisionAdapter{
		handler:     handler,
		nextHandler: nextHandler,
		typeOf:      typeOf,
	}
}

func (c *CollisionAdapter) Handling(spritePositions []sprite.ISprite, from, to int) {
	fromSprite := spritePositions[from].(sprite.ISprite)
	toSprite := spritePositions[to].(sprite.ISprite)
	if spritePositions[to] == nil {
		// TODO:討論  可以放這  或放world
		spritePositions[to] = spritePositions[from]
		spritePositions[from] = nil
		return
	}
	if reflect.TypeOf(fromSprite).String() == string(c.typeOf) {
		c.Collision(fromSprite, toSprite, spritePositions)
		spritePositions[from] = nil
	} else {
		c.next(c.nextHandler, spritePositions, from, to)
	}
}

func (c *CollisionAdapter) Collision(fromSprite, toSprite sprite.ISprite, spritePositions []sprite.ISprite) bool {
	fmt.Println(fmt.Sprintf("--%v >> %v", reflect.TypeOf(fromSprite), reflect.TypeOf(toSprite)))
	isDead := c.handler.Collision(fromSprite, toSprite, spritePositions)
	if isDead {
		toSprite.Remove(&spritePositions)
	}
	return isDead
}

func (c *CollisionAdapter) next(handler CollisionHandler, spritePositions []sprite.ISprite, from, to int) {
	if handler != nil {
		handler.Handle(spritePositions, from, to)
	}
}
