package domain

import (
	"fmt"
)

type PassValidator struct {
	IValidator
	nextHandler IValidator
}

func NewPassValidator(next IValidator) IValidator {
	return &PassValidator{
		nextHandler: next,
	}
}

func (p *PassValidator) Validation(entity ValidatorEntity) int {
	if len(entity.KeyIndex) == 1 && entity.KeyIndex[0] == "-1" {
		player := entity.Player
		if player.GetBig2().TopPlayer != player {
			fmt.Println(fmt.Sprintf("玩家 %s PASS, 頂牌為 %s", player.GetName(), player.GetBig2().TopPlay.ShowCard()))
			return -1
		} else {
			fmt.Println(fmt.Sprintf("你不能在新的回合中喊 PASS"))
			return 0
		}
	}
	return p.nextHandler.Validation(entity)
}
