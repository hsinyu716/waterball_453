package player

import (
	"bufio"
	"cosmos.cards.showdown/internal/common"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type Human struct {
	common.PlayerAdapter
}

func (h *Human) NameHimself(i int) string {
	return fmt.Sprintf("name %d", i)
}

func (h *Human) TakeTurn() *common.TurnMove {
	var ex *common.ExchangeHands
	if h.HasUsedExchangeHands() {
		ex = nil
	} else {
		ex = h.MakeExchangeHandsDecision()
	}
	turnMove := common.NewTurnMove(&h.PlayerAdapter, ex, nil)
	if h.HasUsedExchangeHands() {
		h.GetExchangeHands().Countdown()
	}
	if ex != nil {
		h.SetExchangeHands(ex)
	}
	return turnMove
}

func (h *Human) MakeExchangeHandsDecision() *common.ExchangeHands {
	fmt.Print("Would you like to perform hands exchange? (y/n): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		h.MakeExchangeHandsDecision()
	}

	input = strings.TrimSuffix(input, "\n")
	if input == "y" {
		selectPlayer := h.FilterOtherPlayer()
		return h.SelectExchangeHandsTarget(selectPlayer)
	} else {
		return nil
	}
}
