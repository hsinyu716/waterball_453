package world

import (
	"cosmos.collision/internal/common/handler"
	"cosmos.collision/internal/common/sprite"
	"cosmos.collision/internal/utils"
	"fmt"
	"reflect"
)

type World struct{}

var (
	spritePositions []sprite.ISprite
)

func (w *World) Init() {
	positions := []int{1, 4, 5, 8, 11, 12, 15, 17, 19, 23, 24, 25, 26}
	spritePositions = make([]sprite.ISprite, 30)

	for _, pos := range positions {
		var s sprite.ISprite
		switch pos {
		case 1, 19, 24:
			s = sprite.NewWater(pos)
		case 5, 8, 12:
			s = sprite.NewHero(pos)
		case 11, 17, 23, 26:
			s = sprite.NewFire(pos)
		case 4, 15, 25:
			s = sprite.NewIce(pos)
		}
		spritePositions[pos] = s
	}
	fmt.Println(spritePositions)
}

func (w *World) Move(from int, to int) {
	//hero := &Handler.HeroHandler{}
	//water := &Handler.WaterHandler{}
	//fire := &Handler.FireHandler{}
	//hero.SetHandler(water)
	//water.SetHandler(fire)
	//fire.SetHandler(nil)

	// CoR
	collisionHandler := handler.NewHeroHandler(handler.NewIceHandler(handler.NewFireHandler(handler.NewWaterHandler(nil))))

	fmt.Println(spritePositions[from])
	fmt.Println(spritePositions[to])
	if spritePositions[from] == nil {
		utils.MsgPrint(utils.DataNil)
		return
	}
	if reflect.TypeOf(spritePositions[from]) == reflect.TypeOf(spritePositions[to]) {
		utils.MsgPrint(utils.DataSameType)
		return
	}
	if spritePositions[to] == nil {
		// TODO:討論是在這層判斷是否進到Handle 還是進Handle再判斷  可以放這  或放Handle
		spritePositions[to] = spritePositions[from]
		spritePositions[from] = nil
	} else {
		collisionHandler.Handle(spritePositions, from, to)
	}
	fmt.Println(from, spritePositions[from])
	fmt.Println(to, spritePositions[to])
	fmt.Println(spritePositions)
}
