package domain

import (
	"cosmos.big2/internal/common/pattern"
	"cosmos.big2/internal/common/poker"
	"fmt"
)

type Validator struct {
	handler     IValidator
	nextHandler IValidator
}

type ValidatorEntity struct {
	Hand     *Hand
	KeyIndex []string
	Cards    []*poker.Card
	Player   IPlayer
	TopPlay  pattern.ICardPattern
}

func (v *ValidatorEntity) SetCards(cards []*poker.Card) {
	v.Cards = cards
}

func (v *ValidatorEntity) GetCards() []*poker.Card {
	return v.Cards
}

type IValidator interface {
	Validation(entity ValidatorEntity) int
}

func (v *Validator) Validation(entity ValidatorEntity) int {
	// 0 該次驗證不過
	// 1 輪下一個驗證
	// -1 pass?
	validation := v.Validation(entity)
	fmt.Println("main")
	if validation == 1 {
		return v.next(entity)
	}
	return validation
}

func (v *Validator) next(entity ValidatorEntity) int {
	if v.nextHandler != nil {
		return v.nextHandler.Validation(entity)
	}
	return 1
}
