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
	spritesMap map[int]interface{}
)

func (w *World) Init() {
	spritePositions := []int{1, 5, 8, 11, 12, 17, 19, 23, 24, 26}
	spritesMap = make(map[int]interface{}, 30)

	for _, pos := range spritePositions {
		var s sprite.ISprite
		switch pos {
		case 1, 19, 24:
			s = sprite.NewWater(pos)
		case 5, 8, 12:
			s = sprite.NewHero(pos)
		case 11, 17, 23, 26:
			s = sprite.NewFire(pos)
		}
		spritesMap[pos] = s
	}
	fmt.Println(spritesMap)
}

func (w *World) Move(from int, to int) {
	//hero := &Handler.HeroHandler{}
	//water := &Handler.WaterHandler{}
	//fire := &Handler.FireHandler{}
	//hero.SetHandler(water)
	//water.SetHandler(fire)
	//fire.SetHandler(nil)

	// CoR
	collisionHandler := handler.NewHeroHandler(handler.NewWaterHandler(handler.NewFireHandler(nil)))

	fmt.Println(spritesMap[from])
	fmt.Println(spritesMap[to])
	if spritesMap[from] == nil {
		utils.MsgPrint(utils.DataNil)
		return
	}
	if reflect.TypeOf(spritesMap[from]) == reflect.TypeOf(spritesMap[to]) {
		utils.MsgPrint(utils.DataSameType)
		return
	}
	if spritesMap[to] == nil {
		// TODO:討論是在這層判斷是否進到Handle 還是進Handle再判斷  可以放這  或放Handle
		spritesMap[to] = spritesMap[from]
		spritesMap[from] = nil
	} else {
		collisionHandler.Handle(spritesMap, from, to)
	}
	fmt.Println(from, spritesMap[from])
	fmt.Println(to, spritesMap[to])
	fmt.Println(spritesMap)
}
