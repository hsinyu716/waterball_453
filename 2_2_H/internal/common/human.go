package common

import (
	"fmt"
	"strings"
)

type Human struct {
	PlayerAdapter
}

func (h *Human) NameHimself(i int) {
	h.name = fmt.Sprintf("name %d", i)
}

func (h *Human) TakeTurn() *TurnMove {
	var ex *ExchangeHands
	if h.HasUsedExchangeHands() {
		ex = nil
	} else {
		ex = h.MakeExchangeHandsDecision()
	}
	turnMove := NewTurnMove(&h.PlayerAdapter, ex, nil)
	if h.HasUsedExchangeHands() {
		h.GetExchangeHands().Countdown()
	}
	if ex != nil {
		h.SetExchangeHands(ex)
	}
	return turnMove
}

func (h *Human) MakeExchangeHandsDecision() *ExchangeHands {
	fmt.Print("Would you like to perform hands exchange? (y/n): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.MakeExchangeHandsDecision()
	}

	input = strings.TrimSuffix(input, "\n")
	if input == "y" {
		selectPlayer := h.filterOtherPlayer()
		return h.selectExchangeHandsTarget(selectPlayer)
	} else {
		return nil
	}
}
