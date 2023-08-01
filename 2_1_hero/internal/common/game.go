package common

import "fmt"

type Game struct {
	attacker *Hero
	attacked *Hero
}

func NewGame(attacker *Hero, attacked *Hero) *Game {
	return &Game{
		attacker: attacker,
		attacked: attacked,
	}
}

type IGame interface {
	AttackStart()
}

func (g *Game) AttackStart() {
	g.NextRound()
}

func (g *Game) NextRound() {
	g.attacker.Attack(g.attacked)
	if g.attacked.IsDead() {
		fmt.Println(fmt.Sprintf("%s 已陣亡, %s獲勝！", g.attacked.GetName(), g.attacker.GetName()))
	} else {
		g.attacker, g.attacked = g.attacked, g.attacker
		g.NextRound()
	}
}
