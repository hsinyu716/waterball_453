package handler

import (
	"cosmos.collision/internal/common/sprite"
	"fmt"
	"reflect"
)

type ICollisionHandler interface {
	Handle(spritePositions []sprite.ISprite, from, to int)
	match(fromSprite, toSprite sprite.ISprite) bool
	Collision(fromSprite, toSprite sprite.ISprite, spritePositions *[]sprite.ISprite)
}

type CollisionHandler struct {
	handler     ICollisionHandler
	nextHandler ICollisionHandler
}

func NewCollisionHandler(handler, nextHandler ICollisionHandler) *CollisionHandler {
	return &CollisionHandler{
		handler:     handler,
		nextHandler: nextHandler,
	}
}

func (c *CollisionHandler) Handle(spritePositions []sprite.ISprite, from, to int) {
	fromSprite := spritePositions[from]
	toSprite := spritePositions[to]
	if toSprite == nil {
		toSprite = sprite.NewNil(to)
	}
	if c.handler.match(fromSprite, toSprite) {
		fmt.Println(fmt.Sprintf("--%v >> %v", reflect.TypeOf(fromSprite), reflect.TypeOf(toSprite)))
		c.handler.Collision(fromSprite, toSprite, &spritePositions)
	} else {
		c.next(c.nextHandler, spritePositions, from, to)
	}
}

func (c *CollisionHandler) next(handler ICollisionHandler, spritePositions []sprite.ISprite, from, to int) {
	if handler != nil {
		handler.Handle(spritePositions, from, to)
	}
}
