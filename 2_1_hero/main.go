package main

import (
	"cosmos.hero/internal/common"
)

func main() {
	hero1 := common.NewHero("Hero1", 1000, new(common.WaterBall))
	hero2 := common.NewHero("Hero2", 1200, new(common.Earth))
	game := common.NewGame(hero1, hero2)
	game.AttackStart()
}
