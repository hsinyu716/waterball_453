package sprite

import (
	"cosmos.collision/internal/utils"
)

const (
	initBlood = 30
	blood     = 10
)

type Hero struct {
	Sprite
	hp int
}

func NewHero(p int) *Hero {
	h := &Hero{
		Sprite: Sprite{
			sign: "H",
		},
		hp: initBlood,
	}
	h.setPosition(p)
	return h
}

func (h *Hero) Move(toSprite ISprite, spritePositions *[]ISprite) {
	position := toSprite.GetPosition()
	(*spritePositions)[position] = h
	h.Remove(spritePositions)
	(*spritePositions)[position].ChangePosition(position)
}

func (h *Hero) AddHp() {
	h.hp += blood
	utils.MsgPrint(utils.DataHeroStrengthen)
}

func (h *Hero) MinusHp() bool {
	h.hp -= blood
	utils.MsgPrint(utils.DataHeroWeaken)
	return h.isDead()
}

func (h *Hero) isDead() bool {
	dead := h.hp < 1
	if dead {
		utils.MsgPrint(utils.DataHeroDead)
	}
	return dead
}
