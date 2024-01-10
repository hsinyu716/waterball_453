package domain

import (
	"cosmos.big2/internal/common/pattern"
	"cosmos.big2/internal/common/patternHandler"
	"fmt"
)

type PatternValidator struct {
	IValidator
	nextHandler IValidator
}

func NewPatternValidator(next IValidator) IValidator {
	return &PatternValidator{
		nextHandler: next,
	}
}

func (p *PatternValidator) Validation(entity ValidatorEntity) int {
	cardPattern :=
		pattern.NewPatternSingle(
			pattern.NewPatternPair(
				pattern.NewPatternStraight(
					pattern.NewPatternFullHouse(nil))))
	play := cardPattern.Validate(entity.Cards)

	if play == nil {
		fmt.Println("此牌型不合法，請再嘗試一次。")
		return 0
	}

	fmt.Print(fmt.Sprintf("玩家 %s ", entity.Player.GetName()))
	handler :=
		patternHandler.NewSingleHandler(
			patternHandler.NewPairHandler(
				patternHandler.NewStraightHandler(
					patternHandler.NewFullHouseHandler(nil))))
	compare := handler.Handle(play, entity.TopPlay)
	if compare == 1 {
		entity.Hand.PlayCard(entity.Cards)
		if entity.Hand.Size() == 0 {
			entity.Player.GetBig2().Winner = entity.Player
			return 1
		}
		entity.Player.GetBig2().TopPlayer = entity.Player
		entity.Player.GetBig2().TopPlay = play
		fmt.Println(fmt.Sprintf("目前頂牌玩家為 %s, 頂牌為 %s ", entity.Player.GetName(), play.ShowCard()))
	} else if compare == 0 {
		fmt.Println(fmt.Sprintf("牌型比頂牌小, 頂牌為%s", entity.Player.GetBig2().TopPlay.ShowCard()))
	}
	return compare
}
