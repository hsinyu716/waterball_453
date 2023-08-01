package common

type AttackStrategy interface {
	Attack(attacker *Hero, attacked *Hero)
}
