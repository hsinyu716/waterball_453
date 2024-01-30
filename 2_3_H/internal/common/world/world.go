package world

import (
	"cosmos.collision/internal/common/handler"
	"cosmos.collision/internal/common/sprite"
	"cosmos.collision/internal/utils"
	"fmt"
)

type World struct {
}

var (
	spritePositions []sprite.ISprite
	worldLength     = 30
)

func (w *World) Init() {
	positions := []int{1, 4, 5, 8, 11, 12, 15, 17, 19, 23, 24, 25, 26}
	spritePositions = make([]sprite.ISprite, worldLength)

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
	w.printSprite()
}

func (w *World) printSprite() {
	fmt.Println("=========")
	for i, p := range spritePositions {
		fmt.Println(i, p)
	}
	fmt.Println("=========")
}

func (w *World) Move(from int, to int) {
	//hero := &Handler.HeroHandler{}
	//water := &Handler.WaterHandler{}
	//fire := &Handler.FireHandler{}
	//hero.SetHandler(water)
	//water.SetHandler(fire)
	//fire.SetHandler(nil)

	// CoR
	collisionHandler :=
		handler.NewMoveToNilHandler(
			handler.NewSameSpriteHandler(
				handler.NewBothDisappearHandler(
					handler.NewHeroWeakenHandler(
						handler.NewHeroStrengthenHandler(nil)))))

	fmt.Println(spritePositions[from])
	fmt.Println(spritePositions[to])
	if spritePositions[from] == nil {
		utils.MsgPrint(utils.DataNil)
	} else {
		collisionHandler.Handle(spritePositions, from, to)
	}
	fmt.Println(from, spritePositions[from])
	fmt.Println(to, spritePositions[to])
	w.printSprite()
}
