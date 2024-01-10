package domain

import (
	"fmt"
	"strconv"
)

type SameValidator struct {
	IValidator
	nextHandler IValidator
}

func NewSameValidator(next IValidator) IValidator {
	return &SameValidator{
		nextHandler: next,
	}
}

func (s *SameValidator) Validation(entity ValidatorEntity) int {
	sameMap := make(map[int]bool)
	same := false
	for _, index := range entity.KeyIndex {
		if index == "" {
			break
		}
		atoi, _ := strconv.Atoi(index)
		if sameMap[atoi] {
			same = true
			break
		}
		sameMap[atoi] = true
		entity.Cards = append(entity.Cards, entity.Hand.GetCards()[atoi])
	}
	// 不可出同一張牌
	if same {
		fmt.Println("不可出同一張牌")
		return 0
	}
	entity.SetCards(entity.Cards)
	return s.nextHandler.Validation(entity)
}
