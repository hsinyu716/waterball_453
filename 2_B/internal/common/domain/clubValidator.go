package domain

import (
	"cosmos.big2/internal/utils"
	"fmt"
	"strconv"
)

type ClubValidator struct {
	IValidator
	nextHandler IValidator
}

func NewClubValidator(next IValidator) IValidator {
	return &ClubValidator{
		nextHandler: next,
	}
}

func (c *ClubValidator) Validation(entity ValidatorEntity) int {
	fmt.Println("in club")
	clubIndex, hasClub := entity.Hand.HasClub()
	if hasClub {
		exists, _ := utils.InArray(strconv.Itoa(clubIndex), entity.KeyIndex)
		// todo: 判斷出梅花
		if !exists {
			fmt.Println("第一回合請輸入包含梅花3的牌型。")
			return 0
		}
	}
	return c.nextHandler.Validation(entity)
}
