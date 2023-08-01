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
	//fmt.Println(fmt.Sprintf("--- player: %s, card: %s, %d", e.exchanger.GetName(), hand.name, len(hand.cards)))
	//fmt.Println(fmt.Sprintf("___ player: %s, card: %s, %d", e.exchangee.GetName(), e.exchangee.GetHand().name, len(e.exchangee.GetHand().cards)))
	e.exchanger.SetHand(e.exchangee.GetHand())
	e.exchangee.SetHand(hand)
	fmt.Println(fmt.Sprintf("player %s ←→ player %s.", e.exchanger.GetName(), e.exchangee.GetName()))
	//fmt.Println(fmt.Sprintf("--- player: %s, card: %s, %d", e.exchanger.GetName(), e.exchanger.GetHand().name, len(hand.cards)))
	//fmt.Println(fmt.Sprintf("___ player: %s, card: %s, %d", e.exchangee.GetName(), e.exchangee.GetHand().name, len(e.exchangee.GetHand().cards)))
}

func (e *ExchangeHands) Countdown() {
	e.countdown--
	if e.countdown == 0 {
		e.exchange()
	}
}
