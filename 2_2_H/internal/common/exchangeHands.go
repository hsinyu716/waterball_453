package common

import "fmt"

type ExchangeHands struct {
	countdown int64
	exchanger IPlayer
	exchangee IPlayer
}

type IExchangeHands interface {
	Countdown()
}

func NewExchangeHands(exchanger IPlayer, exchangee IPlayer) *ExchangeHands {
	exHands := &ExchangeHands{
		countdown: 3,
		exchanger: exchanger,
		exchangee: exchangee,
	}
	exHands.exchange()
	return exHands
}

func (e *ExchangeHands) exchange() {
	hand := e.exchanger.GetHand()
	e.exchanger.SetHand(e.exchangee.GetHand())
	e.exchangee.SetHand(hand)
	fmt.Println(fmt.Sprintf("player %s ←→ player %s.", e.exchanger.GetName(), e.exchangee.GetName()))
}

func (e *ExchangeHands) Countdown() {
	e.countdown--
	if e.countdown == 0 {
		e.exchange()
	}
}
