package domain

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Human struct {
	Player
}

func NewHuman() IPlayer {
	return &Human{}
}

func (h *Human) NameSelf(name string) {
	if name == "" {
		fmt.Print("Enter text: ")
		_, _ = fmt.Scanf("%s", &name)
	}
	h.name = name
}

func (h *Human) TakeTurn(input string) int {
	h.GetHand().CardList()
	if len(input) == 0 {
		fmt.Print("輸入要打的牌: ")
	}
	reader := bufio.NewReader(os.Stdin)
	if text, _ := reader.ReadString('\n'); true {
		text = strings.Trim(text, " \n")
		var (
			keyIndex []string
		)

		if len(input) > 0 {
			text = input
		}
		keyIndex = strings.Split(text, " ")

		validator := NewPassValidator(NewClubValidator(NewSameValidator(NewPatternValidator(nil))))

		validateEntity := ValidatorEntity{
			Hand:     h.GetHand(),
			KeyIndex: keyIndex,
			Cards:    nil,
			Player:   h,
			TopPlay:  h.GetBig2().TopPlay,
		}
		return validator.Validation(validateEntity)
	}
	return 0
}
