package sprite

import (
	"cosmos.collision/internal/utils"
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
		hp: 30,
	}
	h.setPosition(p)
	return h
}

func (h *Hero) AddHp() {
	h.hp += 10
	utils.MsgPrint(utils.DataHeroWater)
}

func (h *Hero) MinusHp() bool {
	h.hp -= 10
	utils.MsgPrint(utils.DataHeroFire)
	return h.isDead()
}

func (h *Hero) isDead() bool {
	dead := h.hp < 1
	if dead {
		utils.MsgPrint(utils.DataHeroDead)
	}
	return dead
}
