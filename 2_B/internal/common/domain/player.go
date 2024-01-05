package domain

import "fmt"

type Player struct {
	name string
	hand *Hand
	big2 *Big2
}

type IPlayer interface {
	NameSelf(name string)
	TakeTurn(input string) int
	GetHand() *Hand
	SetHand(h *Hand)
	GetName() string
	SetBig2(big2 *Big2)
	GetBig2() *Big2
}

func (p *Player) NameSelf(name string) {
	p.name = name
	p.hand = NewHand(name)
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) TakeTurn(input string) int {
	fmt.Println(fmt.Sprintf("輪到%s了", p.GetName()))
	p.GetHand().CardList()
	return 0
}

func (p *Player) GetHand() *Hand {
	return p.hand
}

func (p *Player) SetHand(h *Hand) {
	p.hand = h
}

func (p *Player) SetBig2(big2 *Big2) {
	p.big2 = big2
}

func (p *Player) GetBig2() *Big2 {
	return p.big2
}
